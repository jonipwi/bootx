# BootX Artifact Inventory

This directory contains supplied compiled artifacts and visual evidence. It intentionally separates opaque binaries from source, governance, and research documentation.

## Boot artifacts

| File | Format | SHA-256 |
|---|---|---|
| [bootx.img](boot/bootx.img) | raw BIOS disk image | `53AD76FDB61F12FBF9167777FDE92A75076695FCED11FF87B8B027E475DBF6AD` |
| [bootx.iso](boot/bootx.iso) | ISO image | `76C751870FCAE1904B25820F1978033A514DF4ED16D430D33C7F89757F30600D` |
| [bootx.vdi](boot/bootx.vdi) | VirtualBox disk image | `45551492AAD3F00EBD24069000E8E4701CF98A3E32F146198155CBA94D299218` |

## Media

- [bootx.jpg](media/bootx.jpg) — screenshot reporting a BootX Kernel v0.1 VirtualBox run with E1000, DHCP, ARP, and ping behavior.

## Assurance limitation

The current checkout does not contain the source and build system required to reproduce or source-audit these images. Hashes identify the audited files; they do not establish security, safety, ownership, or redistribution permission.

Under the [BootX Common Good Ethical Use License](../LICENSE), opaque binary artifacts must not be redistributed until Corresponding Source, ownership, embedded-component rights, and build provenance are established.

See the [evidence audit](../docs/handbook/01-repository-evidence-audit.md), [progress gate](../PROGRESS.md), and [installation guide](../docs/guides/install-guide.md).
