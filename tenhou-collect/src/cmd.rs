use clap::{value_parser, Arg, ArgAction, ArgMatches, Command, Parser, Subcommand};
use std::path::PathBuf;

pub fn new_cmd() -> Command {
    let config_file_flag = Arg::new("config")
        .short('c')
        .long("config")
        .value_name("FILE")
        .global(true)
        .value_parser(value_parser!(PathBuf))
        .help("Provides a config file");

    Command::new("")
        .bin_name("tenhou")
        .version("v1")
        .about("about")
        .arg(config_file_flag)
        .args_conflicts_with_subcommands(true)
        .subcommand_required(true)
        .subcommand(crawl_cmd())
        .subcommand(catalog_cmd())
        .subcommand(log_cmd())
}

pub fn crawl_cmd() -> Command {
    Command::new("crawl").about("crawl tenhou data")
}

pub fn catalog_cmd() -> Command {
    Command::new("catalog")
}

pub fn log_cmd() -> Command {
    Command::new("log").about("log convert")
}
