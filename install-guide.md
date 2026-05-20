# BootX USB Install Guide

This guide explains how to build a BootX image and write it to a USB flash
drive with Rufus on Windows.

BootX currently has three boot artifacts:

```text
build/bootx.img                    BIOS / legacy raw disk image
build/bootx.iso                    BIOS / legacy El Torito ISO, boot-only
build/bootx.vdi                    BIOS / legacy VirtualBox disk image
build/uefi/EFI/BOOT/BOOTX64.EFI    early UEFI app only
```

Use `build/bootx.img` with Rufus if you want to boot the current BootX kernel
and shell. The UEFI app currently only prints UEFI status text and does not
handoff to the kernel yet.

---

## Safety Warning

Writing BootX to a USB flash drive overwrites the selected USB device.

If you choose the wrong disk in Rufus, you can destroy your operating system or
personal files. Confirm the USB drive size, label, and drive letter before
writing.

---

## Step 1: Install Build Tools

Recommended build environment is WSL/Ubuntu.

Install the required tools:

```bash
sudo apt-get update
sudo apt-get install -y nasm make python3
```

Optional, for QEMU testing:

```bash
sudo apt-get install -y qemu-system-x86
```

---

## Step 2: Build the BIOS USB Image

From WSL, go to the project root:

```bash
cd /mnt/c/Job/bootx
```

Build and verify the raw disk image:

```bash
make clean
make verify
```

Expected success message:

```text
BootX image verified.
```

The Rufus image file is:

```text
C:\Job\bootx\build\bootx.img
```

The same file from WSL is:

```text
/mnt/c/Job/bootx/build/bootx.img
```

Current image layout:

```text
sector 0      BIOS stage1 boot sector
sectors 1-4   stage2 loader
sectors 5-132 protected-mode kernel
sector 133+   FAT12 C: volume with README.TXT
```

---

## Step 3: Test in QEMU First

Before writing a USB drive, test the image in QEMU:

```bash
make run
```

Expected boot output:

```text
BootX stage1 loading...
BootX stage2 ready.
Loading kernel...
Entering protected mode...
BootX Kernel v0.1
BOOTX>
```

Useful shell tests:

```text
help
ls C:\
cat C:\README.TXT
```

---

## Step 4: Write with Rufus

1. Plug in the USB flash drive.
2. Open Rufus.
3. In `Device`, choose the USB flash drive.
4. In `Boot selection`, click `SELECT`.
5. Choose:

```text
C:\Job\bootx\build\bootx.img
```

6. If Rufus asks how to write the image, choose:

```text
Write in DD Image mode
```

7. Confirm that the selected device is the USB flash drive.
8. Click `START`.
9. Accept the warning that the USB drive will be erased.
10. Wait for Rufus to finish, then safely eject the USB drive.

Do not choose ISO mode. BootX is a raw disk image, not an ISO installer.

---

## Step 5: Boot the USB Drive

BootX currently requires BIOS or legacy CSM boot for the kernel path.

In firmware settings:

- Enable `Legacy Boot` or `CSM` if available.
- Disable `Secure Boot`.
- Open the one-time boot menu.
- Choose the USB entry that is not marked `UEFI`, if your firmware shows both.

If the computer is UEFI-only and has no CSM/legacy mode, `build/bootx.img` will
not boot the current kernel on that machine.

---

## Optional: Build the Current UEFI App

The UEFI artifact is useful for testing that firmware can launch BootX's UEFI
entry point, but it does not boot the protected-mode kernel yet.

Build it:

```bash
./scripts/build-uefi.sh
```

Output:

```text
build/uefi/EFI/BOOT/BOOTX64.EFI
```

To test it on a FAT32 USB manually:

1. Format a USB drive as FAT32.
2. Create this directory on the USB:

```text
EFI\BOOT
```

3. Copy `BOOTX64.EFI` to:

```text
EFI\BOOT\BOOTX64.EFI
```

4. Boot the USB in UEFI mode.

Expected behavior: it prints BootX UEFI status text and returns to firmware.

---

## Linux Direct Write Alternative

Rufus is recommended on Windows. On native Linux, you can write the raw image
directly with `dd`.

List disks:

```bash
lsblk
```

Find the whole USB disk, for example:

```text
/dev/sdb
```

Do not write to a mounted directory and do not `cd /dev/sdb`; `/dev/sdb` is a
block device, not a directory. Also do not write to a partition like
`/dev/sdb1` for this raw image.

Write the image:

```bash
sudo dd if=build/bootx.img of=/dev/sdX bs=4M conv=fsync status=progress
sync
sudo eject /dev/sdX
```

Replace `/dev/sdX` with the real USB disk, for example `/dev/sdb`.

---

## Current Hardware Limitations

BootX is still experimental.

Known limitations:

- Current kernel boot path is BIOS/legacy only.
- The UEFI app does not hand off to the kernel yet.
- Disk access uses early ATA PIO assumptions.
- FAT support is FAT12, root directory only, DOS 8.3 names only.
- Keyboard support is basic PS/2-style polling.
- Networking is tested in QEMU NE2000 mode and VirtualBox Intel PRO/1000 mode.
- Many modern laptops do not expose legacy BIOS/CSM or PS/2-compatible keyboard behavior.

QEMU remains the recommended test environment:

```bash
make run
make run-net
```

---

## Optional: Build a VirtualBox Disk

Use a VirtualBox hard-disk image for the full BootX shell, FAT C: volume, and
network tests.

Build the raw image first:

```bash
make verify
```

Convert it to VDI:

```bash
make vdi
```

Attach this file to an IDE storage controller as a hard disk:

```text
C:\Job\bootx\build\bootx.vdi
```

BootX's protected-mode FAT code currently reads sectors with ATA PIO from the
legacy primary IDE port. If the VDI is attached to VirtualBox SATA/AHCI, the BIOS
can still boot it, but BootX will print `FAT: unable to mount C:` after startup.

Use legacy BIOS boot mode in VirtualBox, not EFI mode.

For networking in VirtualBox, set the adapter type to:

```text
Intel PRO/1000 MT Desktop (82540EM)
```

Do not use `PCnet-FAST III`; BootX does not include an AMD PCnet driver.

## Optional: Build a Boot-Only VirtualBox ISO

VirtualBox can also boot BootX from an El Torito ISO, but the current protected
mode kernel cannot mount C: from that ISO path. The kernel reads C: with ATA PIO
from LBA 69, which works when BootX is attached as a hard disk, not when BIOS has
booted a CD/floppy-emulated El Torito image.

Install an ISO builder in WSL:

```bash
sudo apt-get update
sudo apt-get install -y xorriso
```

Build the ISO:

```bash
make iso
```

Attach this file to the VirtualBox optical drive:

```text
C:\Job\bootx\build\bootx.iso
```

Use legacy BIOS boot mode in VirtualBox, not EFI mode.

---

## Recover the USB Drive

To reuse the USB drive normally, repartition and format it with Windows Disk
Management, Rufus, or your operating system's disk tool.

On Linux:

```bash
sudo fdisk /dev/sdX
sudo mkfs.vfat -F 32 /dev/sdX1
```

Be certain `/dev/sdX` is the USB flash drive before formatting.
