use log::error;
use sysinfo::Components;

#[repr(C)]
#[derive(Debug)]
pub enum SensorError {
    Success = 0,
    NoComponents = 1,
    NoTemperature = 2,
}

#[repr(C)]
pub struct SensorResult {
    pub temperature: f32,
    pub error: SensorError,
}

#[no_mangle]
pub extern "C" fn sensors() -> SensorResult {
    let components = match Components::new_with_refreshed_list() {
        Ok(components) => components,
        Err(err) => {
            error!("Failed to initialize components: {}", err);
            return SensorResult {
                temperature: 0.0,
                error: SensorError::NoComponents,
            };
        }
    };

    if let Some(component) = (&components).into_iter().next() {
        if let Some(temp) = component.temperature() {
            SensorResult {
                temperature: temp,
                error: SensorError::Success,
            }
        } else {
            error!("Unknown temperature: {}", component.label());
            SensorResult {
                temperature: 0.0,
                error: SensorError::NoTemperature,
            }
        }
    } else {
        error!("No components found");
        SensorResult {
            temperature: 0.0,
            error: SensorError::NoComponents,
        }
    }
}

#[cfg(test)]
pub mod test {
    use super::*;

    #[test]
    fn simulated_main_function() {
        let result = sensors();

        match result.error {
            SensorError::Success => {
                println!("Temperature reading successful: {} °C", result.temperature);
                assert!(result.temperature >= 0.0); // Basic sanity check
            }
            SensorError::NoComponents => {
                println!("No components found");
                assert_eq!(result.temperature, 0.0);
            }
            SensorError::NoTemperature => {
                println!("No temperature reading available");
                assert_eq!(result.temperature, 0.0);
            }
        }

        // Ensure we got some kind of valid result
        assert!(matches!(
            result.error,
            SensorError::Success | SensorError::NoComponents | SensorError::NoTemperature
        ));
    }
}
