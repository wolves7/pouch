# PouchContainer: How to Provide Container Service Based on OCI

*By Allen Sun, Alibaba Group*

PouchContainer is an open source project which aims at helping shape container standard. It encapsulates OCI runtime and provides container management for users as a container engine, so that container technology could become the foundation for applications in the Cloud era. Currently PouchContainer team is busy preparing 1.0.0 GA release. 

Historically, PouchContainer is one fundamental system software in Alibaba's infrastructure. It helps to process all online transactions smoothly on millions of containers. At present it is used to cover applicable scenarios beyond online business. 

To become a general container engine for every scenario in production, PouchContainer seeks the way which supports several OCI-compatible container runtimes. This action made container service totally out of box:

* runc: container runtime based on Linux cgroups and namespaces;
* katacontainers: container runtime based on hypervisor;
* runlxc: container runtime based on LXC especially on legacy kernels;

## Architecture Based on OCI

OCI-compatible runtimes are the key components in PouchContainer's Architecture. 

![Ecosystem Architecture](static_files/pouch_ecosystem_architecture_no_logo.png)

In the middle right part of architecture, three OCI-compatible runtimes are listed there. 

## OCI Runtime Specification

Mostly business are running in containers based on runC. Lxcfs and advanced features such as diskquota are used to harden isolation among containers and between host and container. Enhanced runC containers plays quite huge role in Alibaba's data centers. 

There are some extra-ordinary applications which care much about security, then containers provided by OCI-compatible runtime katacontainers are the ones users prefer. In addition, those applications, who are close to system and kernel configuration, also take advantages of katacontainers to provide a brand new guest kernel. It is worth mentioning that PouchContainer also made it to run a legacy kernel such as 2.6.32 in katacontainer's guest OS. With this effort, users can containerize legacy applications which depend on legacy kernels and make use of image to speed up delivery.

Legacy issue is always bothering, but still there to be taken care of. runlxc is also an OCI-compatible runtime based on LXC to be compatible with minority of applications on kernel 2.6.32. runlxc has not open sourced yet.

OCI provides a standard of container runtime for industry. Based on OCI's powerful ability, PouchContainer builds a container engine which interacts with end-users easily and smoothly. 

## Kubernetes Native

PouchContainer tries to support all the specific container runtimes in Kubernetes. To achieve this, PouchContainer implements CRI by itself. The CRI implementation passes container runtimes and enhancement of them to Kubernetes in order to build absolutely reliable container service in production.

## Learn More about PouchContainer

PouchContainer also brings other fantastic features to end-users. Want to learn more? Please feel free to visit it on [PouchContainer GitHub](https://github.com/alibaba/pouch).






