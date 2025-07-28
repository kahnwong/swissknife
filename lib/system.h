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
