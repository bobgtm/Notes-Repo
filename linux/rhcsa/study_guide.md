# RHCSA (EX200) Study Plan — Aug 2026 → Exam by 31 Dec 2026

**Target sit date: week of 14 December 2026.** That leaves a two-week buffer before your deadline for a retake or a reschedule.

Total runway: 17 study weeks (W0 = 20–23 Aug for lab build, W1 starts Mon 24 Aug).

---

## 0. Read this first — the 2026 exam is not the exam most books describe

Red Hat's official EX200 page (last published 1 July 2026) states the exam **is based on Red Hat Enterprise Linux 10**. Most study guides, YouTube courses, and Udemy classes on the market right now are written for RHEL 8 or 9. The differences are not cosmetic:

| Change | What it means for you |
|---|---|
| **Flatpak is now an objective** | "Configure access to Flatpak repositories" and "Install and remove Flatpak software packages" are explicitly listed. No RHEL 8/9 guide covers this. Nearly nobody is teaching it. It is low-hanging points. |
| **Containers are gone** | The entire "Manage containers" category (podman run/start/stop/list, container images, rootless containers as systemd services) that dominated the RHEL 8/9 exam is **not on the current objectives list**. If a course spends three modules on podman, that time is wasted for EX200. |
| **MBR is gone** | Objectives now say "List, create, and delete partitions on **GPT** disks." The old "MBR and GPT" wording is dropped. Learn `parted`/`gdisk` on GPT; don't burn time on `fdisk` MBR quirks. |
| **VDO and Stratis are not listed** | Two storage technologies that were heavily featured in RHEL 8-era prep material do not appear. Skip them. |
| **Tuning profiles are listed** | "Manage tuning profiles" — `tuned-adm`. Small topic, easy points, frequently skipped. |
| **Bootloader modification is listed** | "Modify the system bootloader" — GRUB2 config, `grubby`, kernel arguments, default kernel. |
| **autofs is listed** | Direct and indirect maps. This trips up a lot of candidates. |

**Practical consequence:** buy or borrow a RHEL 9 guide if you must (Sander van Vugt's and Asghar Ghori's are the two standards, and the underlying material is 90% transferable), but treat this plan's objective list as authoritative, and *build your lab on RHEL 10*.

**Get RHEL 10 legally and free:** register at developers.redhat.com for the **Red Hat Developer Subscription for Individuals**. It gives you up to 16 systems at no cost, real `subscription-manager` registration, and access to the actual CDN repos — which matters, because "Install and update software packages from Red Hat Content Delivery Network" is a literal objective. Rocky/AlmaLinux are excellent 95% substitutes but will not let you practice `subscription-manager` or the RHSM repo layout.

### The two rules that decide pass/fail

1. **Everything must survive a reboot, with no intervention.** Red Hat says this explicitly. A candidate who does every task correctly in the running system and forgets `systemctl enable`, `firewall-cmd --permanent`, or an `/etc/fstab` entry fails. **Reboot after every single lab in this plan.** Make it a reflex, not a checklist item.
2. **You cannot bring notes, and there is no internet.** But the shipped documentation *is* available: `man`, `info`, `/usr/share/doc`, and `pinfo`. Your real skill isn't memorization — it's knowing that `/usr/share/doc/pacemaker/` style example configs exist and that `man 5 fstab`, `man semanage-fcontext`, and `/usr/share/doc/autofs/` will hand you the syntax under pressure. **Practice looking things up in man pages, deliberately, every week.**

---

## 1. Lab environment (Week 0 — do this before 24 Aug)

You are building four VMs. Everything in this plan assumes them.

| VM | Role | Specs | Notes |
|---|---|---|---|
| `rhcsa-a` | Primary victim | 2 vCPU, 2 GB RAM, 20 GB OS disk **+ two blank 5 GB disks** | This is the one you break and rebuild. |
| `rhcsa-b` | Second host | 2 vCPU, 2 GB, 20 GB | For SSH, `scp`/`rsync`, NFS client work. |
| `rhcsa-nfs` | NFS/autofs server | 1 vCPU, 1 GB, 20 GB | Exports for the autofs objectives. |
| `rhcsa-gold` | Never touched | 20 GB | Snapshot source. Clone from this. |

Two extra blank disks on `rhcsa-a` are non-negotiable — the storage objectives (GPT partitions, PV/VG/LV, swap, non-destructive extension) need free block devices, and you will destroy them repeatedly.

**Snapshot discipline is the highest-leverage habit in this whole plan.** Before every lab: snapshot. After every lab: revert. On Proxmox:

```bash
qm snapshot 101 clean --description "pristine RHEL10"
qm rollback  101 clean
```

Write yourself a two-line shell function for rollback on day one. If reverting takes more than five seconds of thought, you will start avoiding destructive practice — and destructive practice is the entire point.

**Also install `sos`** (`dnf install sos`) and use `sosreport`/`sos report` once, just to see what a support bundle contains. Not on the exam. Extremely useful for the mental model of "where does state live."

---

## 2. Weekly rhythm

You asked to study as often as possible with reasonable breaks. This is the cadence:

| Day | Time | Structure |
|---|---|---|
| Mon–Fri | 75–90 min | 15 min concept read → 55 min hands-on lab → 10 min notes + flashcards |
| Saturday | 2.5–3 hr | Integration lab: combine the week's topic with everything prior. Reboot-verify. |
| Sunday | **Off**, or 20 min flashcard review only | Real rest. Non-negotiable. |

Inside a weekday session, work 25 min / 5 min break / 25 min / 5 min break / 25 min. The five-minute breaks are load-bearing; skipping them is why people burn out in week 9.

**Deload:** the last weekday of every fourth week (Weeks 4, 8, 12, 16) is review-only — no new material, just re-drilling the weakest items your triage system has flagged.

**The spaced-repetition rule:** every Saturday session begins with 20 minutes of "cold recall" on a topic from **three weeks ago**, done from a blank shell with no notes. This is the single mechanic that separates people who pass from people who studied the same hours and failed. Recognition is not recall; the exam only tests recall.

---

## 3. The 17-week schedule

### W0 · 20–23 Aug — Build the lab
Install RHEL 10 on all four VMs (minimal install, no GUI, root + one regular user). Register with `subscription-manager`. Take the `clean` snapshot on each. Verify you can revert.

---

### W1 · 24–30 Aug — Shell, filesystem, files, and documentation
**Objectives:** Access a shell prompt and issue commands with correct syntax · Create, delete, copy, and move files and directories · Create hard and soft links · Locate, read, and use system documentation

- **Mon** — Filesystem Hierarchy Standard. Walk `/etc`, `/var`, `/usr`, `/opt`, `/srv`, `/run`, `/proc`, `/sys`. For each, say out loud what belongs there. `man hier` is your friend.
- **Tue** — Navigation and file ops: `cd -`, `pushd`/`popd`, `cp -a` vs `cp -r`, `mv`, `rm -rf` (and why `--one-file-system` exists), `mkdir -p`, `rmdir`, brace expansion `touch file{1..10}.txt`, globbing vs regex (they are *different* — internalize this now).
- **Wed** — Links. Create hard links and symlinks. Break a symlink. Observe inode numbers with `ls -li` and `stat`. Prove to yourself a hard link cannot cross filesystems and a symlink can. Delete the target of each and see what happens.
- **Thu** — Documentation drill. `man -k` (`apropos`), `man 5 passwd` vs `man 1 passwd` (**understand man sections — this saves you on exam day**), `info coreutils`, `ls /usr/share/doc/`. Find the answer to "what does the `nofail` mount option do" using only local docs.
- **Fri** — `tar`, `gzip`, `bzip2`, `xz`. Create, list, extract. Extract to a different directory (`-C`). Preserve permissions and SELinux contexts (`--xattrs --selinux`).
- **Sat** — Integration: archive `/etc` with permissions intact, move it to `rhcsa-b`, extract, verify nothing changed. Reboot both. Verify again.

---

### W2 · 31 Aug – 6 Sep — Text processing, redirection, vim
**Objectives:** Use input-output redirection · Use grep and regular expressions to analyze text · Create and edit text files

- **Mon** — Redirection deeply: `>`, `>>`, `2>`, `2>&1`, `&>`, `<`, `<<`, `<<<`, `|`, `tee`, `/dev/null`. **Understand why `command > file 2>&1` works and `command 2>&1 > file` doesn't** — order of evaluation. This is a classic.
- **Tue** — `grep` and BRE/ERE: anchors `^ $`, classes `[[:digit:]]`, quantifiers, `-i -v -c -n -r -E -o -A -B -C`. Practice on `/var/log/` and `/etc/passwd`.
- **Wed** — `sed` basics (`s///g`, `-i`, address ranges, `d`), `awk` basics (`$1`, `NF`, `-F:`, `BEGIN`/`END`). You need *survival level* awk, not mastery.
- **Thu** — `cut`, `sort`, `uniq -c`, `wc`, `head`, `tail -f`, `tr`, `column`, `find` (by name, size, mtime, perm, user; with `-exec` and `-delete`).
- **Fri** — **vim, seriously.** Modes, `:w`, `:wq!`, `:q!`, `dd`, `yy`, `p`, `u`, `Ctrl-r`, `/search`, `:%s/old/new/g`, `:set nu`, visual block mode. You will edit config files under time pressure with no GUI. Do the whole session in vim.
- **Sat** — Integration: write a one-liner pipeline that reports the ten largest files under `/var`, and another that lists every user with a login shell. Reboot. Rerun.

---

### W3 · 7–13 Sep — Users, groups, permissions
**Objectives:** Create, delete, and modify local user accounts · Change passwords and adjust password aging · Create, delete, and modify local groups and group memberships · List, set, and change standard ugo/rwx permissions · Manage default file permissions · Diagnose and correct file permission problems

- **Mon** — `/etc/passwd`, `/etc/shadow`, `/etc/group`, `/etc/gshadow` field by field. `useradd`, `usermod`, `userdel -r`, `groupadd`, `groupmod`, `groupdel`. `/etc/login.defs`, `/etc/skel`, `/etc/default/useradd`.
- **Tue** — UID/GID ranges, system vs regular accounts, `id`, `groups`, primary vs supplementary groups, `usermod -aG` (**and why forgetting `-a` is destructive**).
- **Wed** — Password aging: `chage -l/-M/-m/-W/-I/-E`, `passwd -l/-u/-e/-S`, shadow field mapping. Lock an account three different ways and explain the difference between each.
- **Thu** — Permissions: octal and symbolic, `chmod`, `chown`, `chgrp`. Directory `x` vs `r` (**the most misunderstood thing in Linux permissions**).
- **Fri** — Special bits: setuid, setgid, sticky. `umask` — compute the resulting mode by hand for files vs directories, then verify. Set a persistent umask for a specific user.
- **Sat** — Integration: build a shared group directory (`/srv/projectx`) where group members can collaborate, files inherit the group, and nobody can delete another's files. Reboot. Verify with two test users.

---

### W4 · 14–20 Sep — sudo, SSH, remote access · **deload Friday**
**Objectives:** Access remote systems using SSH · Log in and switch users in multi-user targets · Configure privileged access · Configure key-based authentication for SSH · Securely transfer files between systems

- **Mon** — `su` vs `su -` vs `sudo -i` vs `sudo -s`. Login vs non-login shells and which files each sources (`/etc/profile`, `~/.bash_profile`, `~/.bashrc`, `/etc/bashrc`).
- **Tue** — sudo: `/etc/sudoers` syntax, **always edit with `visudo`**, drop-in files in `/etc/sudoers.d/`, `%wheel`, `NOPASSWD:`, command aliases, restricting to specific commands.
- **Wed** — SSH: `ssh`, `ssh -p`, `~/.ssh/config`, `ssh-keygen`, `ssh-copy-id`, `authorized_keys`, permissions on `~/.ssh` (`700`) and `authorized_keys` (`600`) — get these wrong and it silently fails.
- **Thu** — `sshd_config` hardening: `PermitRootLogin`, `PasswordAuthentication`. `scp`, `sftp`, `rsync -avz`. Transfer between all three VMs.
- **Fri** — **Deload.** Re-drill W1–W3 weak spots only.
- **Sat** — Integration: passwordless key auth from `rhcsa-a` to `rhcsa-b` for a non-root user, that user gets passwordless sudo for `systemctl` only, password auth disabled. Reboot both. Verify.

---

### W5 · 21–27 Sep — systemd: units, services, targets
**Objectives:** Start, stop, and check the status of network services · Start and stop services and configure services to start automatically at boot · Configure systems to boot into a specific target automatically · Boot systems into different targets manually

- **Mon** — systemd model: units, unit types (`.service`, `.target`, `.socket`, `.timer`, `.mount`), `systemctl list-units`, `list-unit-files`.
- **Tue** — `start`, `stop`, `restart`, `reload`, `enable`, `disable`, `mask`, `unmask`, `status`, `is-enabled`, `is-active`. **`enable` ≠ `start`. `mask` ≠ `disable`.** Know exactly what `mask` does and why it exists.
- **Wed** — Unit file anatomy: `[Unit]`, `[Service]`, `[Install]`, `Type=`, `ExecStart=`, `Restart=`, `WantedBy=`. Unit file precedence: `/etc/systemd/system` beats `/run/systemd/system` beats `/usr/lib/systemd/system`. `systemctl edit` and drop-ins. `daemon-reload`.
- **Thu** — Targets: `multi-user.target`, `graphical.target`, `rescue.target`, `emergency.target`. `systemctl get-default`, `set-default`, `isolate`.
- **Fri** — Write your own service unit from scratch that runs a script, enable it, reboot, confirm it started.
- **Sat** — Integration: mixed scenario — set default target, mask a service, write a custom unit, break one deliberately and read `systemctl status` output to diagnose. Reboot-verify all.

---

### W6 · 28 Sep – 4 Oct — Boot process, GRUB, recovery, tuned
**Objectives:** Boot, reboot, and shut down a system normally · Interrupt the boot process in order to gain access to a system · Modify the system bootloader · Manage tuning profiles

- **Mon** — Boot chain: firmware → GRUB2 → kernel + initramfs → systemd → default target. Know each handoff.
- **Tue** — GRUB2: `/etc/default/grub`, `grub2-mkconfig`, BIOS vs UEFI config paths, `grubby --info=ALL`, `grubby --update-kernel`, setting the default kernel, adding a persistent kernel argument.
- **Wed** — **Root password recovery.** `rd.break` method: interrupt GRUB, append `rd.break`, `mount -o remount,rw /sysroot`, `chroot /sysroot`, `passwd`, `touch /.autorelabel`, exit twice. **Do this five times until it's muscle memory.** It is on the exam more often than not.
- **Thu** — Emergency and rescue targets. Recover from a broken `/etc/fstab` entry (this is the #1 self-inflicted "my VM won't boot" scenario, and also a very common exam task).
- **Fri** — `tuned-adm list`, `active`, `profile`, `recommend`. Create a custom profile in `/etc/tuned/`. Small topic, cheap points.
- **Sat** — Integration: break your own system three ways (lost root password, bad fstab, wrong default target), then recover each. This is the most valuable single session in the entire plan.

---

### W7 · 5–11 Oct — Processes, scheduling priority, logs and journals
**Objectives:** Identify CPU/memory intensive processes and kill processes · Adjust process scheduling · Locate and interpret system log files and journals · Preserve system journals

- **Mon** — `ps aux`, `ps -ef`, `ps -eo`, `top`, `uptime`, `free`, `pgrep`, `pidof`. Job control: `&`, `jobs`, `fg`, `bg`, `Ctrl-z`, `nohup`.
- **Tue** — Signals: `kill`, `killall`, `pkill`, SIGTERM(15) vs SIGKILL(9) vs SIGHUP(1). **Understand why SIGKILL is the last resort, not the first.**
- **Wed** — Priority: nice values −20..19, `nice`, `renice`, why only root can lower a nice value. Read the `NI` and `PR` columns in `top` correctly.
- **Thu** — `journalctl`: `-u`, `-p`, `-b`, `-b -1`, `--since`/`--until`, `-f`, `-x`, `-n`, `--disk-usage`, `--vacuum-time`.
- **Fri** — **Persistent journals:** `mkdir -p /var/log/journal`, `Storage=persistent` in `/etc/systemd/journald.conf`, restart, reboot, then confirm `journalctl -b -1` returns output. That last verification *is* the exam task. Also cover `rsyslog` and `/etc/rsyslog.conf` facilities/priorities, and `logrotate`.
- **Sat** — Integration: generate load, find and kill the offender, renice a process, prove journal persistence survives two reboots.

---

### W8 · 12–18 Oct — Software: dnf, rpm, repos, Flatpak · **deload Friday**
**Objectives:** Configure access to RPM repositories · Install and remove RPM software packages · Install and update software packages from Red Hat CDN, a remote repository, or from the local file system · Configure access to Flatpak repositories · Install and remove Flatpak software packages

- **Mon** — `dnf install/remove/update/search/info/provides/list/history`. `dnf provides */filename` — extremely useful. `dnf history undo`.
- **Tue** — Repo config: `/etc/yum.repos.d/*.repo` fields (`[id]`, `name`, `baseurl`, `enabled`, `gpgcheck`, `gpgkey`). Point at a local directory (`file:///`), an HTTP mirror, and the CDN. `dnf config-manager`. `createrepo_c` for a local repo.
- **Wed** — `subscription-manager register/attach/repos --list/--enable`. Module streams: `dnf module list/enable/install/reset`.
- **Thu** — `rpm -qa`, `-qi`, `-ql`, `-qf`, `-qc`, `-qd`, `-V` (verify — read the output codes), `-ivh`, `--nodeps` (and why you shouldn't).
- **Fri** — **Deload**, but spend 30 min on **Flatpak**: `flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo`, `flatpak remotes`, `flatpak search`, `flatpak install`, `flatpak list`, `flatpak uninstall`, `flatpak remote-delete`, and the `--user` vs system-wide distinction. New objective, almost nobody teaches it, cheap points.
- **Sat** — Integration: build a local repo from downloaded RPMs, configure a client to use it, install from it, add a Flatpak remote and install a package. Reboot. Verify both persist.

---

### W9 · 19–25 Oct — Networking
**Objectives:** Configure IPv4 and IPv6 addresses · Configure hostname resolution · Configure network services to start automatically at boot

- **Mon** — NetworkManager model. **The critical distinction: a *device* is hardware; a *connection profile* is config. `nmcli device` vs `nmcli connection`.** Profiles live in `/etc/NetworkManager/system-connections/` as keyfiles. The old `/etc/sysconfig/network-scripts/ifcfg-*` files are gone.
- **Tue** — `nmcli con show`, `add`, `mod`, `up`, `down`, `delete`, `reload`. Set a static IPv4: address, prefix, gateway, DNS, `ipv4.method manual`, `connection.autoconnect yes`.
- **Wed** — IPv6: `ipv6.method`, `ipv6.addresses`, link-local vs global. `ip addr`, `ip route`, `ip -6`. `nmtui` as the safety net when nmcli syntax escapes you under pressure — **know it exists**.
- **Thu** — Hostname resolution: `hostnamectl set-hostname`, `/etc/hosts`, `/etc/resolv.conf` (and why NetworkManager owns it), `/etc/nsswitch.conf`, `getent hosts`, `dig`, `host`, `ss -tulpn`.
- **Fri** — Time: `timedatectl`, `chronyd`, `/etc/chrony.conf`, `chronyc sources`. Set timezone and an NTP server.
- **Sat** — Integration: static IPv4 + IPv6 + DNS + hostname on `rhcsa-a`, reachable from `rhcsa-b` by name. **Reboot and verify** — networking config that doesn't survive reboot is the classic zero-point answer.

---

### W10 · 26 Oct – 1 Nov — firewalld
**Objectives:** Restrict network access using firewalld and firewall-cmd · Configure firewall settings using firewall-cmd/firewalld

- **Mon** — Zone model: `public`, `trusted`, `internal`, `drop`, `block`, `work`, `home`. Default zone, interface-to-zone binding, source-to-zone binding.
- **Tue** — **Runtime vs permanent.** `firewall-cmd --add-service=http` (runtime only, lost on reload) vs `--permanent` (config only, not live) vs `--permanent ... && --reload`. Also `--runtime-to-permanent`. **This single concept accounts for a huge share of lost firewall points.**
- **Wed** — Services vs ports: `--add-service`, `--add-port=8080/tcp`, `--list-all`, `--get-services`, `--info-service`.
- **Thu** — Rich rules and `--add-source`, `--remove-*`, `--set-default-zone`, `--change-interface`.
- **Fri** — Port forwarding and masquerade (`--add-forward-port`, `--add-masquerade`).
- **Sat** — Integration: install httpd, serve on a non-standard port, open it in firewalld permanently, verify from `rhcsa-b`. This will fail on SELinux — **note the failure and leave it broken.** You'll fix it in W13, and the memory of the failure is what makes SELinux stick.

---

### W11 · 2–8 Nov — Storage I: partitions, filesystems, fstab, swap
**Objectives:** List, create, and delete partitions on GPT disks · Configure systems to mount file systems at boot by UUID or label · Create, mount, unmount, and use VFAT, ext4, and XFS file systems · Add new partitions and swap to a system non-destructively

- **Mon** — Block device landscape: `lsblk`, `blkid`, `fdisk -l`, `df -h`, `du -sh`, `/dev/disk/by-uuid/`, `/dev/disk/by-id/`.
- **Tue** — GPT partitioning with `parted`: `mklabel gpt`, `mkpart`, `print`, `rm`, `set`. Also `gdisk`. `partprobe` / `udevadm settle` after changes.
- **Wed** — Filesystems: `mkfs.xfs`, `mkfs.ext4`, `mkfs.vfat`. Labels: `xfs_admin -L`, `e2label`, `fatlabel`. `xfs_growfs` vs `resize2fs` (**XFS can grow but never shrink — know this cold**).
- **Thu** — `/etc/fstab`: all six fields. Mount by UUID *and* by LABEL. **Always `mount -a` before rebooting**, and use `nofail` while practicing so a typo doesn't strand you in emergency mode. `systemctl daemon-reload` after editing fstab.
- **Fri** — Swap: `mkswap`, `swapon`, `swapoff`, `swapon --show`, `free -h`, swap priority, persistent swap partition and swap file in fstab.
- **Sat** — Integration: on the blank 5 GB disk, create two GPT partitions — one XFS mounted by UUID at `/data`, one swap with priority set. Reboot. `lsblk`, `swapon --show`, `df -h` must all confirm.

---

### W12 · 9–15 Nov — Storage II: LVM, NFS, autofs · **deload Friday**
**Objectives:** Create and remove physical volumes · Assign physical volumes to volume groups · Create and delete logical volumes · Extend existing logical volumes · Add new logical volumes non-destructively · Mount and unmount network file systems using NFS · Configure autofs

- **Mon** — LVM layering: PV → VG → LV → filesystem. `pvcreate`, `pvs`, `pvdisplay`, `vgcreate`, `vgs`, `vgextend`, `lvcreate -L`/`-l`, `lvs`, `lvdisplay`.
- **Tue** — Extending: `lvextend -L +2G -r /dev/vg/lv` (**the `-r` flag resizes the filesystem in the same step — learn it and you'll never forget the second command**). Do it the manual two-step way too, so you understand what `-r` is doing. Extend a VG onto a second PV first.
- **Wed** — `lvreduce` on ext4 (unmount → `e2fsck -f` → `resize2fs` → `lvreduce`), and why this is impossible on XFS. Remove LVs, VGs, PVs cleanly. LVM snapshots.
- **Thu** — NFS: export from `rhcsa-nfs` (`/etc/exports`, `exportfs -rav`, `nfs-server`, firewall), mount from `rhcsa-a` manually, then persistently in fstab with `_netdev`.
- **Fri** — **Deload**, plus 40 min on **autofs**: `/etc/auto.master`, `/etc/auto.master.d/*.autofs`, indirect maps (`/shares  /etc/auto.shares`), direct maps (`/-  /etc/auto.direct`), wildcard maps (`*  server:/export/&`). This is the objective candidates most often fumble. Get the direct-vs-indirect distinction crisp.
- **Sat** — Integration: full storage scenario — new PV, extend VG, create LV, XFS, fstab by UUID, extend it live, plus an autofs-mounted home directory from the NFS server. Reboot. Everything must come back.

---

### W13 · 16–22 Nov — SELinux
**Objectives:** Set enforcing and permissive modes · List and identify SELinux file and process context · Restore default file contexts · Manage SELinux port labels · Use Boolean settings to modify system SELinux settings

- **Mon** — The mental model (see Deep Dives doc). `getenforce`, `setenforce 0/1`, `/etc/selinux/config`, `sestatus`. Understand that permissive still *logs*.
- **Tue** — Contexts: `ls -Z`, `ps -Z`, `id -Z`. The `user:role:type:level` format — for RHCSA, **the type is what matters** (`httpd_sys_content_t`, `ssh_home_t`, etc.).
- **Wed** — `restorecon -Rv`, `chcon` (**and why `chcon` is a trap — it does not survive a relabel**), `semanage fcontext -a -t TYPE "/path(/.*)?"` followed by `restorecon`. This two-command pattern is the single most-tested SELinux skill.
- **Thu** — Port labels: `semanage port -l`, `semanage port -a -t http_port_t -p tcp 8404`. **Go back and fix the W10 httpd-on-a-weird-port lab.**
- **Fri** — Booleans: `getsebool -a`, `setsebool -P` (**the `-P` is the whole exam question**), `semanage boolean -l`.
- **Sat** — Integration + troubleshooting: `ausearch -m AVC -ts recent`, `sealert`, reading `/var/log/audit/audit.log`. Break three services with wrong contexts and diagnose each from the audit log alone. Then `touch /.autorelabel` and reboot to see a full relabel.

---

### W14 · 23–29 Nov — Shell scripting, at/cron/timers · *(light week — Thanksgiving)*
**Objectives:** Conditionally execute code · Use looping constructs · Process script inputs ($1, $2) · Process output of shell commands within a script · Schedule tasks using at, cron, and systemd timer units

- **Mon** — Script anatomy: shebang, `chmod +x`, exit codes, `$?`, `$0 $1 $@ $# $*`, quoting rules (**`"$@"` vs `$*` — know the difference**), `$(command)`.
- **Tue** — Conditionals: `if/elif/else/fi`, `test`, `[ ]`, `[[ ]]`, string vs numeric comparison operators (`-eq` vs `==`), file tests (`-f -d -e -r -w -x -z -n`), `&&`/`||`, `case`.
- **Wed** — Loops: `for i in ...`, `for` over command output, `while read line`, `until`, `break`, `continue`. Read a file line by line.
- **Thu** — `at`, `atq`, `atrm`. `cron`: `crontab -e/-l/-r/-u`, `/etc/crontab`, `/etc/cron.d/`, the five time fields, `/etc/cron.{daily,hourly}`. `cron.allow`/`cron.deny`.
- **Fri** — **systemd timers:** `.timer` + matching `.service`, `OnCalendar=`, `OnBootSec=`, `Persistent=`, `systemctl list-timers`, `systemd-analyze calendar "Mon *-*-* 09:00:00"`. Enable and verify across a reboot.
- **Sat** *(or skip to Sun if traveling)* — Integration: write a script taking two arguments, validating them, looping over a directory, and logging output — then schedule it three ways (at, cron, timer).

---

### W15 · 30 Nov – 6 Dec — Mock exam 1 + remediation
- **Mon** — **Mock Exam 1.** Fresh VM from `rhcsa-gold`. Three hours. Timer running. No notes, no internet, only man pages. Use the lab scenario bundle (`rhcsa-labs/`).
- **Tue** — Grade honestly. Every failed task goes on a written weak-list.
- **Wed–Fri** — Drill the weak-list, one topic per day, hardest first.
- **Sat** — Re-run only the tasks you failed, cold.

**By this point, book your exam.** Red Hat individual exams book up; schedule for the week of 14 Dec now. Remote exam requires a compatible machine and a clear room — test your setup *this week*, not the night before.

---

### W16 · 7–13 Dec — Mock exams 2 and 3 · **deload Friday**
- **Mon** — Mock Exam 2 (different scenario set). Three hours.
- **Tue** — Grade + drill.
- **Wed** — Mock Exam 3. Three hours.
- **Thu** — Grade + drill.
- **Fri** — **Deload.** Light review only. Rest matters more than one more lab.
- **Sat** — Speed drills: root password reset, LVM extend, SELinux context fix, firewalld permanent rule, fstab by UUID, persistent journal. Each under five minutes. Repeat until they are automatic.

---

### W17 · 14–20 Dec — Sit the exam
- **Mon** — Final speed drills (as above). Nothing new.
- **Tue** — Rest. Sleep.
- **Wed/Thu** — **Exam day.**
- **Fri** — Buffer / decompression.

---

## 4. Exam-day tactics

1. **Read every question before starting.** Some tasks depend on others; some are 30 seconds. Harvest the cheap ones first.
2. **Set the root password task first** if it appears — you can't do anything else without access.
3. **Reboot early, around the halfway mark**, not just at the end. If something breaks a reboot, you want to find out with 90 minutes left, not 5.
4. **Reboot again at the end** and re-verify every task. Budget 15 minutes for this. It is not optional.
5. When stuck on syntax, go to `man` immediately. Do not burn four minutes guessing. `man -k <topic>` then read the EXAMPLES section at the bottom — most man pages have one.
6. **Do not leave a task blank.** Partial credit exists on most tasks.
7. If you break the system badly, remember rescue mode and `rd.break` are available to you during the exam too.

---

## 5. Resources

- **Official objectives** (authoritative, check monthly for changes): redhat.com EX200 page
- **Red Hat docs**, free: docs.redhat.com — the RHEL 10 "Configuring and managing" series maps almost 1:1 to the objectives
- **Sander van Vugt**, *RHCSA Cert Guide* — best structured book; get the newest edition available
- **Asghar Ghori**, *RHCSA Study Guide* — more labs, good complement
- **Red Hat Learning Subscription** — expensive, but includes RH124/RH134 and exam attempts; worth pricing out against the standalone exam fee
- **`man` pages and `/usr/share/doc`** — genuinely your best resource, because they're the only one you'll have on exam day

---

## 6. Tracking

The `rhcsa-triage/` system emails you each morning with the day's tasks and flags anything overdue. The `rhcsa-labs/` system generates broken systems and grades your fixes. Both are in the accompanying directories, and both are themselves RHCSA practice — you'll deploy them with systemd timers, shell scripts, and cron.

