use clap::Parser;
use std::path::PathBuf;

#[derive(Parser)]
#[command(version)]
#[command(about = "", long_about = None)]
pub struct CliFlag {
    /// Sets a custom config file
    #[arg(short, long, value_name = "FILE", default_value = "config.yaml")]
    pub config: Option<PathBuf>,

    /// Turn debugging information on
    #[arg(long)]
    pub debug: bool,
}
