mod cmd;
mod flag;

use crate::flag::CliFlag;
use clap::Parser;
use clap::{Arg, ArgAction, Command};
use tracing::info;

fn main() {
    // let cli = CliFlag::parse();
    let cmd = Command::new("")
        .bin_name("tenhou")
        .version("v1")
        .about("about")
        .subcommand_required(true)
        .subcommand(Command::new("crawl"))
        .subcommand(Command::new("catalog"))
        .subcommand(Command::new("log").about("log convert"));

    tracing_subscriber::fmt::init();

    let _matches = cmd.get_matches();
    // info!("Hello, world!,{},{:?}", cli.debug, cmd.get_long_flag());
    info!("Hello, world!");
}
