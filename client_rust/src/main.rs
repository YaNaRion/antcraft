use raylib::prelude::*;

struct Window {
    rl: RaylibHandle,
    title: String,
}

impl Window {
    fn borro(rl: RaylibHandle, title: String) -> Window {
        Self {
            rl: rl,
            title: title,
        }
    }
}

struct Unit;
fn main() {
    let title = "Hello, Raylib!";
    let (rl, thread) = raylib::init().size(800, 450).title(title).build();

    let mut window = Window::borro(rl, title.to_string());

    while !window.rl.window_should_close() {
        let mut d = window.rl.begin_drawing(&thread);
        d.clear_background(Color::RAYWHITE);
        d.draw_text("Hello, Raylib in Rust!", 190, 200, 20, Color::BLACK);
    }
}
