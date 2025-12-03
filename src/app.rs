use crate::device::{get_available_devices, Device};
use crate::settings::AppSettings;
use color_eyre::eyre;
use colors_transform::{self, Color as TColor};
use crossterm::event::{self, Event, KeyCode};
use ratatui::{
    layout::{Constraint, Direction, Layout},
    style::{Color, Modifier, Style},
    symbols::DOT,
    text::{Line, Span},
    widgets::{block::Padding, Block, List, ListItem, ListState, Paragraph},
    DefaultTerminal, Frame,
};

pub struct App {
    should_exit: bool,
    settings: AppSettings,
    colors: AppColors,
    list: DeviceList,
}

struct AppColors {
    title_font: Color,
    title_background: Color,
    selected_device: Color,
}

struct DeviceList {
    devices: Vec<Device>,
    state: ListState,
}

impl Default for App {
    fn default() -> Self {
        let settings = match AppSettings::load() {
            Ok(s) => s,
            Err(e) => panic!("{}", e),
        };

        let devices = match get_available_devices(&settings.mount.point, &settings.mount.options) {
            Ok(d) => d,
            Err(e) => panic!("{}", e),
        };

        Self {
            should_exit: false,
            colors: AppColors {
                title_font: hex_to_rgb(&settings.colors.title_font),
                title_background: hex_to_rgb(&settings.colors.title_background),
                selected_device: hex_to_rgb(&settings.colors.selected_device),
            },
            list: DeviceList {
                devices: devices,
                state: ListState::default(),
            },
            settings: settings,
        }
    }
}

impl App {
    pub fn run(&mut self, mut terminal: DefaultTerminal) -> eyre::Result<()> {
        while !self.should_exit {
            terminal.draw(|frame| render(frame, self))?;
            if let Event::Key(key) = event::read()? {
                self.handle_key(key.code)
            }
        }

        Ok(())
    }

    fn handle_key(&mut self, key: KeyCode) {
        match key {
            KeyCode::Enter => self.current_device_toggle_mount(),
            KeyCode::Char('q') => self.should_exit = true,
            KeyCode::Char('r') => self.reload_device_list(),
            KeyCode::Down | KeyCode::Char('j') => self.list.state.select_next(),
            KeyCode::Up | KeyCode::Char('k') => self.list.state.select_previous(),
            _ => {}
        }
    }

    fn reload_device_list(&mut self) {
        let devices =
            match get_available_devices(&self.settings.mount.point, &self.settings.mount.options) {
                Ok(d) => d,
                Err(e) => panic!("{}", e),
            };

        self.list.devices = devices;
    }

    fn current_device_toggle_mount(&mut self) {
        if self.list.devices.is_empty() {
            return;
        }

        let current_index = match self.list.state.selected() {
            Some(u) => u,
            None => return,
        };

        match self.list.devices[current_index].toggle_mount() {
            Ok(_) => {}
            Err(e) => panic!("{}", e),
        }
    }
}

fn render(frame: &mut Frame, app: &mut App) {
    let outer_block = Block::new().padding(Padding::proportional(1));
    let inner_block = Block::new()
        .title(Line::raw(" Available MTP Devices "))
        .title_style(
            Style::new()
                .fg(app.colors.title_font)
                .bg(app.colors.title_background),
        )
        .padding(Padding::proportional(1));

    let inner_area = outer_block.inner(frame.area());
    frame.render_widget(&inner_block, inner_area);

    let chunks = Layout::default()
        .direction(Direction::Vertical)
        .constraints([Constraint::Min(0), Constraint::Length(1)])
        .split(inner_block.inner(inner_area));

    if app.list.devices.is_empty() {
        let msg =
            Line::raw("No MTP devices attached.").style(Style::new().fg(Color::Rgb(60, 60, 60)));
        frame.render_widget(msg, chunks[0]);
    } else {
        let items = app
            .list
            .devices
            .iter()
            .map(|x| x.to_string().into())
            .collect::<Vec<ListItem>>();
        app.list.state.select(Some(0));

        let list = List::new(items).highlight_style(Style::new().fg(app.colors.selected_device));
        frame.render_stateful_widget(list, chunks[0], &mut app.list.state);
    }

    let help_message = Line::from(vec![
        Span::styled("q ", Style::default().add_modifier(Modifier::BOLD)),
        Span::raw("quit "),
        Span::raw(DOT),
        Span::styled(" r ", Style::default().add_modifier(Modifier::BOLD)),
        Span::raw("refresh "),
        Span::raw(DOT),
        Span::styled(" 󰌑 ", Style::default().add_modifier(Modifier::BOLD)),
        Span::raw("mount/umount"),
    ])
    .style(Style::new().fg(Color::Rgb(60, 60, 60)));

    let footer = Paragraph::new(help_message);
    frame.render_widget(footer, chunks[1]);
}

fn hex_to_rgb(s: &String) -> Color {
    let rgb = match colors_transform::Rgb::from_hex_str(s) {
        Ok(c) => c,
        Err(e) => panic!("{}", e.message),
    };

    let (r, g, b) = rgb.as_tuple();
    let r = r as u8;
    let g = g as u8;
    let b = b as u8;
    Color::Rgb(r, g, b)
}
