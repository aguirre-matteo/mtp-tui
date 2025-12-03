#[derive(Debug)]
pub enum AppError {
    DeviceMounted(String),
    DeviceNotMounted(String),
    FailedToExecuteCommand(String),
    JmtpfsFailed(String),
    Utf8ConversionFailed(String),
    JmtpfsWrongFormat(String),
    MountpointExists(String),
    FailedToCreateMountpoint(String),
    MountpointNotFound(String),
    FuseFailed(String),
    FailedToRemoveMountpoint(String),
}

impl std::fmt::Display for AppError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            AppError::DeviceMounted(s) => {
                write!(f, "Tried to mount device {} but it's already mounted", s)
            }
            AppError::DeviceNotMounted(s) => {
                write!(f, "Tried to umount device {} but it's not mounted", s)
            }
            AppError::FailedToExecuteCommand(s) => {
                write!(f, "Failed to execute the command '{}'", s)
            }
            AppError::JmtpfsFailed(s) => write!(f, "jmptfs:\n{}", s),
            AppError::Utf8ConversionFailed(s) => {
                write!(f, "Failed to convert Vec<u8> to UTF8: {}", s)
            }
            AppError::JmtpfsWrongFormat(s) => write!(f, "Wrong jmtpfs's output format: {}", s),
            AppError::MountpointExists(s) => write!(f, "Mountpoint exists: {}", s),
            AppError::FailedToCreateMountpoint(s) => write!(f, "Failed to create: {}", s),
            AppError::MountpointNotFound(s) => write!(f, "Mountpoint not found: {}", s),
            AppError::FuseFailed(s) => write!(f, "fusermount:\n{}", s),
            AppError::FailedToRemoveMountpoint(s) => {
                write!(f, "Failed to remove mountpoint: {}", s)
            }
        }
    }
}
