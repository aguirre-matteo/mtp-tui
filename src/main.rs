mod app;
mod device;
mod errors;
mod settings;

fn main() {
    if let Err(e) = color_eyre::install() {
        panic!("{}", e);
    }

    let terminal = ratatui::init();
    let result = app::App::default().run(terminal);
    ratatui::restore();

    if let Err(e) = result {
        panic!("{}", e);
    }
}
