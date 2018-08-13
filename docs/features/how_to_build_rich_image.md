# How to Build Rich Image

Rich container is a very useful container mode when containerizing 
applications. This mode helps technical staff to complete packaging rich
applications almost with no effort. For more details, please refer to
[Rich Container](./pouch_with_rich_container.md).

To enjoy the ability of rich container, it is very essential to build an image
which supports rich container mode. We call it **rich image**.

## Rich Container Workflow

Before diving deep into rich image's content, a brief review of rich
container's architecture is necessary.

![pouch_with_rich_container](../static_files/pouch_with_rich_container.png)

From the architecture above, we can tell that in the lifecycle of rich
container, there will be six running parts:

* prestart hook (optional)
* systemd(required)
* system services (optional)
* ENTRYPOINT/CMD(required)
* customized agent (optional)
* poststop hook (optional)

Therefore, when building a rich image, six ensential software above should be
covered.

## Building a Rich Image

Rich Image is totally compatible with Docker image and OCI image. So, users
have ability to build rich images based on docker images. What's more, rich
image is totally compatible with Dockerfile.

### Systemd

Systemd is the fundamental software for rich container. Checking the existence
of `systemd` or `/sbin/init` is the first thing to do. 

### System Service

System service is booted by `/sbin/init` or `systemd`. For all system services
which needs installed in rich image, they should be recorded in Dockerfile.
For example, to add a `sshd` system service, a command of 
`ADD sshd.service /etc/systemd/system/sshd.service` should be written in
Dockerfile and the content of this file is like below:

```
[Unit]
Description=OpenBSD Secure Shell server
After=network.target auditd.service
ConditionPathExists=!/etc/ssh/sshd_not_to_be_run

[Service]
EnvironmentFile=-/etc/default/ssh
ExecStart=/usr/sbin/sshd -D $SSHD_OPTS
ExecReload=/bin/kill -HUP $MAINPID
KillMode=process
Restart=on-failure
RestartPreventExitStatus=255
Type=notify

[Install]
WantedBy=multi-user.target
Alias=sshd.service
```


### Customized Agent

### Hooks


### Simple Sample


