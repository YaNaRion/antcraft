use raylib::prelude::*;

struct Window {
    rl: RaylibHandle,
    title: String,
}

impl Window {
<<<<<<< HEAD
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
=======
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
>>>>>>> 64fdb68e278de4d7faac8aae7e5988bd60756730

    while !window.rl.window_should_close() {
        let mut d = window.rl.begin_drawing(&thread);
        d.clear_background(Color::RAYWHITE);
<<<<<<< HEAD
        d.draw_text(text_to_draw, 190, 200, 20, Color::BLACK);
=======
        d.draw_text("Hello, Raylib in Rust!", 190, 200, 20, Color::BLACK);
>>>>>>> 64fdb68e278de4d7faac8aae7e5988bd60756730
    }
}
