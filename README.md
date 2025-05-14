PREVIEW.gif

Mtp-Tui is a terminal user interface for easily mounting and umounting your MTP devices.
Android phones, digital cameras and media players fall under this category.

# Table of Contents
- [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Configuration](#configuration)
- [Nix](#nix)
    - [Flake](#flake)
    - [Modules](#modules)
- [Known issues](#known-issues)
- [Contributing](#contributing)
- [Licence](#licence)

# Getting Started
Mtp-Tui is a terminal user interface for easily mounting and umounting your MTP devices.
Android phones, digital cameras and media players fall under this category.

## Installation
First, clone this repository and run `go build` on its root directory:

```shell
git clone https://github.com/aguirre-matteo/mtp-tui
cd mtp-tui
go build
```

Then copy the binary to somewhere in your $PATH. For example, copy it to `/usr/bin/`:

```shell
cp ./mtp-tui /usr/bin/
```

## Usage 
The `mtp-tui` command must be run with `sudo`, so it can access the
device's storage. Use the `-u` flag for specifying under which user's 
config the program will be running:

```shell
mtp-tui -u yourusername
```

This will open the interface, showing all the detected devices in that moment.
You can move through the elements using the arrow keys or HJKL.

If you haven't already connected your device, it's now time to do so.
In the case of an Android phone, connect it to your computer, swipe down
the screen and you'll find a notification saying that the device has been connected.

<img src="./screenshots/usb-notification.png" width="300" alt="USB Notification">

Then click on it and the USB's Connection details will open. Select the "File Transfer"
option in the "Use USB for" section. This way you'll be able to access your phone's 
storage.

<img src="./screenshots/usb-settings.png" width="300" alt="USB Settings">

Now, in your computer, press "r" on the app and the device list will update. Make sure
you update the list after changing to "File Transfer" mode. Otherwise, you won't be able 
to access your files.

## Configuration
A simple YAML config file can be placed on `/etc/mtp-tui.yml` or `~/.config/mtp-tui.yml`
to configure mounting options. Here's an example config showing all the available options,
and its default values:

```yaml
mount:
  point: /mtp/                             # Where devices will be mounted. If the user is different from root, it will be ~/mtp/
  options: default_permissions,allow_other # Mount options. These ensure your user has access to the drive.
                                           # default_permissions ensures the mounted FS inherits his parent directory's permissions.
```

# Nix
> [!NOTE]
> This guide assumes you have flakes enabled in your Nix config.

To install Mtp-Tui on NixOS, this project provides both a NixOS and a Home-Manager module.

## Flake
First, add this repository to your flake's inputs:

```flake.nix
{
  inputs = {
    # ...
    mtp-tui.url = "github:aguirre-matteo/mtp-tui";
  };
}
```

Then add the NixOS and Home-Manager modules in your configuration. You should also
add this flake's overlay. Otherwise Nix wouldn't be able to find the package:

```flake.nix
{
  outputs = { self, nixpkgs, home-manager, ...}@inputs: {
    nixosConfigurations.nixos = nixpkgs.lib.nixosSystem {
      # ...
      modules = [
        ./configuration.nix
        inputs.mtp-tui.nixosModule # <--- This installs the NixOS module
        
        home-manager.nixosModules.home-manager
        {
          home-manager = {
            # ...
            sharedModules = [
              inputs.mtp-tui.homeManagerModule # <--- This installs the Home-Manager module
            ];
          };
        }
        {
          nixpkgs.overlays = [
            inputs.mtp-tui.overlay # <--- This installs the Nixpkgs overlay
          ];
        }
      ];
    };
  };
}
```

## Modules
These are the options implemented by both NixOS and Home-Manager modules at `programs.mtp-tui`:

| Option     | NixOS module    | Home-Manager module |
|------------|-----------------|---------------------|
| `enable`   | Enables the app | Enable the app      |
| `package`  | The package     | The package         |
| `settings` | /etc config     | ~/.config config    |

Here's an example config:
```configuration.nix
{ pkgs, ... }:

{
  programs.mtp-tui = {
    enable = true;
    package = pkgs.mtp-tui.override { buildGoModule = my.custom.function };
    settings = {
      mount = {
        point = "/home/matteo/Documents/mtp";
        options = "default_permissions,allow_other";
      };
    };
  };
}
```

# Known issues
At the date of writing this guide, there're is no easy way to mount a MTP device without root permissions.
There are also some bugs regarding to LIBMTP and Jmtpfs, and this app throws an exception when trying to
run it without `sudo`, so the best option for now is to use `sudo` along the `-u` flag.

# Contributing
If you want to contribute to this project, you have some options to do so:

- Solving an open issue.
- Open an issue reporting a bug or requesting a feature.
- Spread this project in forums, or recommending it to someone else.

# Licence
This program is distributed under the MIT Licence. See the details [here](LICENSE).
