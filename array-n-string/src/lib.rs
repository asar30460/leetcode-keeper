mod merge_sorted_array_88;

use std::error::Error;
use std::io;

const DIVIDER: &str = "═══════════════════════════════════════════";

fn clear_screen() {
    print!("\x1B[2J\x1B[1;1H");
}

fn print_header(title: &str) {
    println!("\n╔{}╗", DIVIDER);
    println!("║{:^43}║", title);
    println!("╠{}╣", DIVIDER);
}

fn print_footer() {
    println!("╚{}╝\n", DIVIDER);
}

fn print_menu_item(number: &str, description: &str) {
    println!("║  [{}] {:<37}║", number, description);
}

fn print_instruction(instruction: &str) {
    println!("║ {} {:<40}║", "→", instruction);
}

pub fn route() -> Result<(), Box<dyn Error>> {
    clear_screen();
    print_header("ARRAY AND STRING");
    println!("║                                           ║");
    println!("║  Select a question:                       ║");
    println!("║                                           ║");
    print_menu_item("1", "88. Merge Sorted Array");
    println!("║                                           ║");
    print_menu_item("0", "Back to Main Menu");
    println!("║                                           ║");
    print_instruction("Enter your choice and press Enter:");
    print_footer();

    let mut buf = String::new();
    match io::stdin().read_line(&mut buf) {
        Ok(_) => match buf.trim() {
            "1" => merge_sorted_array_88::test()?,
            "0" => return Ok(()),
            _ => println!(
                "\n❌ Invalid input: '{}'. Please enter a valid option.",
                buf.trim()
            ),
        },
        Err(e) => println!("\n❌ Error reading input: {}", e),
    }

    println!("\nPress Enter to continue...");
    io::stdin().read_line(&mut String::new()).unwrap();

    Ok(())
}
