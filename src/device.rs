use crate::errors::AppError;
use std::fmt;
use std::fs::{create_dir_all, remove_dir};
use std::path::Path;
use std::process::Command;

pub struct Device {
    bus: String,
    id: String,
    name: String,
    mounted: bool,
    mountpoint: String,
    mount_options: String,
}

impl Device {
    fn mount(&mut self) -> Result<(), AppError> {
        if self.mounted {
            return Err(AppError::DeviceMounted(format!(
                "{} ({}/{})",
                self.name, self.bus, self.id
            )));
        }

        if Path::new(&self.mountpoint).is_dir() {
            return Err(AppError::MountpointExists(self.mountpoint.clone()));
        }

        if let Err(_) = create_dir_all(&self.mountpoint) {
            return Err(AppError::FailedToCreateMountpoint(self.mountpoint.clone()));
        }

        let devflag = format!("-device={},{}", &self.bus, &self.id);
        let mut cmd = Command::new("jmtpfs");
        cmd.arg(devflag);

        if self.mount_options != "None" {
            cmd.arg(&self.mount_options);
        }

        cmd.arg(&self.mountpoint);
        let cmd = match cmd.output() {
            Ok(o) => o,
            Err(_) => {
                return Err(AppError::FailedToExecuteCommand(format!(
                    "jmtpfs -device={},{}",
                    &self.bus, &self.id
                )));
            }
        };

        if !cmd.status.success() {
            let stderr = match String::from_utf8(cmd.stderr) {
                Ok(o) => o,
                Err(_) => {
                    return Err(AppError::Utf8ConversionFailed(String::from(
                        "'jmtpfs -l' stderr",
                    )));
                }
            };
            return Err(AppError::JmtpfsFailed(stderr));
        }

        self.mounted = true;
        Ok(())
    }

    fn umount(&mut self) -> Result<(), AppError> {
        if !self.mounted {
            return Err(AppError::DeviceNotMounted(format!(
                "{} ({}/{})",
                self.name, self.bus, self.id
            )));
        }

        if !Path::new(&self.mountpoint).is_dir() {
            return Err(AppError::MountpointNotFound(self.mountpoint.clone()));
        }

        let cmd = match Command::new("fusermount")
            .arg("-u")
            .arg(&self.mountpoint)
            .output()
        {
            Ok(o) => o,
            Err(_) => {
                return Err(AppError::FailedToExecuteCommand(format!(
                    "fusermount -u {}",
                    self.mountpoint
                )));
            }
        };

        if !cmd.status.success() {
            let stderr = match String::from_utf8(cmd.stderr) {
                Ok(o) => o,
                Err(_) => {
                    return Err(AppError::Utf8ConversionFailed(format!(
                        "'fusermount -u {}'",
                        self.mountpoint
                    )));
                }
            };
            return Err(AppError::FuseFailed(stderr));
        }

        if let Err(_) = remove_dir(&self.mountpoint) {
            return Err(AppError::FailedToRemoveMountpoint(self.mountpoint.clone()));
        }

        self.mounted = false;
        Ok(())
    }

    pub fn toggle_mount(&mut self) -> Result<(), AppError> {
        if self.mounted {
            self.umount()
        } else {
            self.mount()
        }
    }
}

impl fmt::Display for Device {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let mounted_msg = if self.mounted {
            format!("| Mounted at: {}", self.mountpoint)
        } else {
            String::new()
        };

        write!(
            f,
            "{}\nBus ID: {} | Device ID: {} {}",
            self.name, self.bus, self.id, mounted_msg
        )
    }
}

pub fn get_available_devices(
    base_mountpoint: &String,
    mount_options: &String,
) -> Result<Vec<Device>, AppError> {
    let cmd = match Command::new("jmtpfs").arg("-l").output() {
        Ok(o) => o,
        Err(_) => {
            return Err(AppError::FailedToExecuteCommand(String::from("jmtpfs -l")));
        }
    };

    if !cmd.status.success() {
        let stderr = match String::from_utf8(cmd.stderr) {
            Ok(o) => o,
            Err(_) => {
                return Err(AppError::Utf8ConversionFailed(String::from(
                    "'jmtpfs -l' stderr",
                )));
            }
        };
        return Err(AppError::JmtpfsFailed(stderr));
    }

    let output = match String::from_utf8(cmd.stdout) {
        Ok(o) => o,
        Err(_) => {
            return Err(AppError::Utf8ConversionFailed(String::from(
                "'jmtpfs -l' stdout",
            )));
        }
    };

    let lines_to_skip = match output.lines().position(|x| {
        x == "Available devices (busLocation, devNum, productId, vendorId, product, vendor):"
    }) {
        Some(i) => i + 1,
        None => unreachable!(),
    };

    let mut devices: Vec<Device> = vec![];
    for line in output.lines().skip(lines_to_skip) {
        let fields: Vec<_> = line.split(", ").collect();
        if fields.len() != 6 {
            return Err(AppError::JmtpfsWrongFormat(String::from(line)));
        }

        let mountpoint = format!("{}/{}_{}", base_mountpoint, fields[0], fields[1]);
        let mounted = Path::new(&mountpoint).is_dir();
        let dev = Device {
            bus: String::from(fields[0]),
            id: String::from(fields[1]),
            name: String::from(fields[4]),
            mounted: mounted,
            mountpoint: mountpoint,
            mount_options: mount_options.clone(),
        };
        devices.push(dev);
    }
    Ok(devices)
}
