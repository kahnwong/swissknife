// battery
enum BatteryError {
    BATTERY_SUCCESS = 0,
    BATTERY_NO_BATTERY = 1,
    BATTERY_NO_CYCLE_COUNT = 2,
    BATTERY_MANAGER_ERROR = 3
};

struct BatteryResult {
    unsigned int cycle_count;
    enum BatteryError error;
};

struct BatteryResult battery_cycle_count(void);


// sensors
enum SensorError {
    SENSOR_SUCCESS = 0,
    SENSOR_NO_COMPONENTS = 1,
    SENSOR_NO_TEMPERATURE = 2
};

struct SensorResult {
    float temperature;
    enum SensorError error;
};

struct SensorResult sensors(void);
