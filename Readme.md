# CHroot-BIN

Easily execute chrooted binnary from current environment without chroot command

For example your distroless application chroot dir is `/apps/bash/`

To execute your application, normally you can use bellow command

```bash
chroot /apps/bash/ /usr/bin/bash
```

With Chroot-bin, instead of long command, you can create link to chroot
and execute directly.

```bash
ln -s /usr/bin/chroot-bin /usr/bin/bash # One time

bash
```

## Config

By default, application looks "/apps/"+$(basename $0) path for chdir.
You can change default dir by setting `CHROOTBIN_DIR_DEFAULT` variable
EX. CHROOTBIN_DIR_DEFAULT="/myapps/

Currently two option is avaible for per binnary.
One of them is indicates chrootdir of the application and second one is
binnary path **inside** the chroot dir.

This helpfull if the multiple apps is under same dir.

```bash
CHROOTBIN_DIR="bash=/apps/debian;apt=/apps/debian" ./bash
```

If your link name and bin name is different, you can set custom bin path

```bash
 export CHROOTBIN_DIR="ash=/apps/alpine;apk=/apps/alpine;sh=/apps/alpine" CHROOTBIN_PATH="sh=/bin/ash"
 ln -s chroot-bin sh
 ./sh
 sudo -E ./apk update
```
