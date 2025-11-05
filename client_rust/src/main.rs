use raylib::prelude::*;

struct Window {
    rl: RaylibHandle,
    title: String,
}

impl Window {
    // fn borro(rl: RaylibHandle, title: String) -> Window {
    //     Self {
    //         rl: rl,
    //         title: title,
    //     }
    // }
}

fn main() {
    let window_title = "Hello, Raylib!";
    let (rl, thread) = raylib::init().size(800, 450).title(window_title).build();

    let mut window = Window {
        rl: rl,
        title: window_title.to_string(),
    };

    let text_to_draw = "Hello, Raylib in Rust!";

    while !window.rl.window_should_close() {
        let mut d = window.rl.begin_drawing(&thread);
        d.clear_background(Color::RAYWHITE);
        d.draw_text(text_to_draw, 190, 200, 20, Color::BLACK);
    }
}
