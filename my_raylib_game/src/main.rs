use raylib::prelude::*;

struct Window {
    rl: RaylibHandle,
    title: String,
}

impl Window {
    fn new(rl: RaylibHandle, title: String) -> Self {
        Self {
            rl: rl,
            title: title,
        }
    }
}

struct Unit;
fn main() {
    let title = "Hello, Raylib!";
    let (mut rl, thread) = raylib::init().size(800, 450).title(title).build();

    let window = Window::new(rl, title.to_string());

    while !rl.window_should_close() {
        let mut d = rl.begin_drawing(&thread);
        d.clear_background(Color::RAYWHITE);
        d.draw_text("Hello, Raylib in Rust!", 190, 200, 20, Color::BLACK);
    }
}
