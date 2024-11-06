#!/bin/bash
set -e

log() {
  local message="$1"
  echo "$(date +'%Y-%m-%d %H:%M:%S') - $message"
}

RESOURCE=${1:-"$HOME/resource"}
KUBERNETES_VERSION=${2:-"v1.31.2"}
CONTAINERD_VERSION=${3:-"v1.7.23"}
RUNC_VERSION=${4:-"v1.2.0"}

if [ ! -d "$RESOURCE" ] || [ ! -r "$RESOURCE" ]; then
  log "Error: RESOURCE directory $RESOURCE does not exist or is not readable"
  exit 1
fi

ARCH=$(uname -m)
case $ARCH in
aarch64)
  ARCH="arm64"
  ;;
x86_64)
  ARCH="amd64"
  ;;
*)
  log "Error: Unsupported architecture $ARCH. Supported architectures are: aarch64, x86_64"
  exit 1
  ;;
esac

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
if [[ "$OS" != "linux" ]]; then
  log "Error: Unsupported OS $OS"
  exit 1
fi

if ss -tulpn | grep -q ":6443"; then
  log "Error: Port 6443 is already in use. Please free up the port and try again."
  exit 1
fi

function install_kubernetes_software() {
  kubernetesPath="$RESOURCE/$ARCH/kubernetes/$KUBERNETES_VERSION"

  if [ ! -d "$kubernetesPath" ] || [ ! -r "$kubernetesPath" ]; then
    log "Error: Directory $kubernetesPath does not exist or is not readable"
    exit 1
  fi

  if [ ! -f "$kubernetesPath/kubeadm" ]; then
    log "Error: File $kubernetesPath/kubeadm does not exist"
    exit 1
  fi

  if ! install -m 755 "$kubernetesPath/kubeadm" /usr/local/bin/kubeadm; then
    log "Error: Failed to install kubeadm"
    exit 1
  fi

  if [ ! -f "$kubernetesPath/kubelet" ]; then
    log "Error: File $kubernetesPath/kubelet does not exist"
    exit 1
  fi

  if ! install -m 755 "$kubernetesPath/kubelet" /usr/local/bin/kubelet; then
    log "Error: Failed to install kubelet"
    exit 1
  fi

}

containerdService=$(
  cat <<EOF
[Unit]
Description=containerd container runtime
Documentation=https://containerd.io
After=network.target local-fs.target dbus.service

[Service]
ExecStartPre=-/sbin/modprobe overlay
ExecStart=/usr/local/bin/containerd

Type=notify
Delegate=yes
KillMode=process
Restart=always
RestartSec=5

LimitNPROC=infinity
LimitCORE=infinity

# Comment TasksMax if your systemd version does not supports it.
# Only systemd 226 and above support this version.
TasksMax=infinity
OOMScoreAdjust=-999

[Install]
WantedBy=multi-user.target
EOF
)

function install_containerd() {
  log "install containerd..."

  containerdPath="$RESOURCE/$ARCH/containerd/$CONTAINERD_VERSION/"
  if [ ! -d "$containerdPath" ] || [ ! -r "$containerdPath" ]; then
    log "Error: Directory $containerdPath does not exist or is not readable"
    exit 1
  fi

  chmod -R 755 "${containerdPath}/bin/"

  cp -r "${containerdPath}/bin/" /usr/local/bin/

  if ! ctr --version; then
    log "Error: Failed to start containerd service"
    exit 1
  fi

  pause_image="aliyun.io/pause:3.10"

  containerd config default | sed -e "s/SystemdCgroup: false/SystemdCgroup: true/g" -e "s/sandbox_image = .*/sandbox_image = \"$pause_image\"/g" | tee /etc/containerd/config.toml

  if ! echo "$containerdService" | tee /usr/lib/systemd/system/containerd.service >/dev/null; then
    log "Error: Failed to write to /usr/lib/systemd/system/containerd.service"
    exit 1
  fi

  if ! systemctl daemon-reload; then
    log "Error: Failed to reload systemd daemon"
    exit 1
  fi

  if ! systemctl enable --now containerd; then
    log "Error: Failed to start containerd service"
    exit 1
  fi

  log "install runc..."
  runcPath="$RESOURCE/$ARCH/runc/$RUNC_VERSION/"
  if [ ! -d "$runcPath" ] || [ ! -r "$runcPath" ]; then
    log "Error: Directory $runcPath does not exist or is not readable"
    exit 1
  fi

  install -m 755 "$runcPath/runc" /usr/local/bin/runc
}

install_kubernetes_software

if systemctl is-active --quiet containerd; then
  log "containerd is already running, skipping installation."
else
  log "containerd is not running, proceeding with installation."
  install_containerd
fi

log "kubernetes software installation completed successfully."

exit 0
