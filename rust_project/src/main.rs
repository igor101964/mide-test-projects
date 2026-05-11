fn greet(name: &str) -> String {
    format!("Hello, {}!", name)
}

fn main() {
    println!("Rust Project test");
    println!("  {}", greet("mide"));
}
