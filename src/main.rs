use std::{error::Error, io};

const DIVIDER: &str = "═══════════════════════════════════════════";
const SECTION_DIVIDER: &str = "───────────────────────────────────────────";

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

fn main() -> Result<(), Box<dyn Error>> {
    loop {
        clear_screen();
        print_header("LEETCODE KEEPER");
        println!("║                                           ║");
        println!("║  Select a category:                       ║");
        println!("║                                           ║");
        print_menu_item("1", "Array And String");
        println!("║                                           ║");
        print_menu_item("0", "Exit");
        println!("║                                           ║");
        print_instruction("Enter your choice and press Enter:");
        print_footer();

        let mut buf = String::new();
        match io::stdin().read_line(&mut buf) {
            Ok(_) => match buf.trim() {
                "1" => arrnstr::route()?,
                "0" => {
                    println!("Thank you for using LeetCode Keeper!");
                    break;
                }
                _ => {
                    println!(
                        "\n❌ Invalid input: '{}'. Please enter a valid option.",
                        buf.trim()
                    );
                    pause();
                }
            },
            Err(e) => {
                println!("\n❌ Error reading input: {}", e);
                pause();
            }
        }
    }

    Ok(())
}

fn pause() {
    println!("\nPress Enter to continue...");
    io::stdin().read_line(&mut String::new()).unwrap();
}
