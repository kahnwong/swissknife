extern crate log;
extern crate sysinfo;

use log::error;
use sysinfo::{Components, System};


#[no_mangle]
pub extern "C" fn sensors() -> f32 {
    let mut system = System::new_all();
    system.refresh_all();
    let components = Components::new_with_refreshed_list();

    let mut output: f32 = 0.0;
    if let Some(component) = (&components).into_iter().next() {
        if let Some(temperature) = component.temperature() {
            output = temperature;
        } else {
            error!("Unknown temperature: {}", component.label())
        }
    }

    output
}


#[cfg(test)]
pub mod test {
    use super::*;

    #[test]
    fn simulated_main_function() {
        sensors();
    }
}
