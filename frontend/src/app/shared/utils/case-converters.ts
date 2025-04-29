import {
    Facilities,
    PropertySurroundings,
} from '../../models/accommodation.model';

export function facilitiesToSnakeCase(facilities: Facilities) {
    return {
        wifi: facilities.wifi,
        air_condition: facilities.airCondition,
        tv: facilities.tv,
    };
}

export function propertySurroundsToSnakeCase(
    propertySurrounds: PropertySurroundings
) {
    return {
        restaurant: propertySurrounds.restaurant,
        bar: propertySurrounds.bar,
    };
}
