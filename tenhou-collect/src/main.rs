mod cmd;
mod config;

use std::fs;
use tracing::info;

fn main() {
    let cmd = cmd::new_cmd();
    tracing_subscriber::fmt::init();
    let matches = cmd.get_matches();

    match matches.subcommand() {
        Some(("log", sub_m)) => {
            let config = sub_m
                .get_one::<std::path::PathBuf>("config")
                .expect("`port`is required");
            info!("log sub m: {:?}, config {:?}", sub_m, config);

            info!("current path: {:?}", std::env::current_dir());
            let content = fs::read_to_string(config).expect("read config path");
            let config_data: config::Config =
                serde_yaml::from_str(content.as_str()).expect("serialize config failed");

            info!("config: {:?}", config_data)
        }
        _ => {}
    }
}
