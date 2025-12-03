use config::{Config, ConfigError, File, FileFormat};
use std::env;

#[derive(serde::Deserialize)]
pub struct MountSettings {
    pub point: String,
    pub options: String,
}

#[derive(serde::Deserialize)]
pub struct ColorsSettings {
    pub title_font: String,
    pub title_background: String,
    pub selected_device: String,
}

#[derive(serde::Deserialize)]
pub struct AppSettings {
    pub mount: MountSettings,
    pub colors: ColorsSettings,
}

impl AppSettings {
    pub fn load() -> Result<Self, ConfigError> {
        let home = match env::var("HOME") {
            Ok(h) => h,
            Err(_) => {
                return Err(ConfigError::Message(String::from(
                    "Couldn't get the $HOME variable",
                )))
            }
        };

        let default_mountpoint = format!("{}/mtp", home);

        let mut settings = Config::builder()
            .set_default("mount.options", "None")?
            .set_default("mount.point", default_mountpoint)?
            .set_default("colors.title_font", "#FFFFFF")?
            .set_default("colors.title_background", "#5F5FD7")?
            .set_default("colors.selected_device", "#C86EC8")?
            .add_source(File::new("/etc/mtp-tui", FileFormat::Yaml).required(false));

        settings = if let Ok(xdg_config_home) = env::var("XDG_CONFIG_HOME") {
            settings.add_source(
                File::new(&format!("{}/mtp-tui", &xdg_config_home), FileFormat::Yaml)
                    .required(false),
            )
        } else {
            settings.add_source(
                File::new(&format!("{}/.config/mtp-tui", &home), FileFormat::Yaml).required(false),
            )
        };

        settings.build()?.try_deserialize()
    }
}
