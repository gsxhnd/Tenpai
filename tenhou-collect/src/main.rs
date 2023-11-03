mod flag;

use crate::flag::CliFlag;
use clap::Parser;
use tracing::info;

fn main() {
    let cli = CliFlag::parse();
    tracing_subscriber::fmt::init();
    info!("Hello, world!,{}", cli.debug);
}
