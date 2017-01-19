FROM scratch

COPY kubeadm /kubeadm

ENTRYPOINT ["/kubeadm"]