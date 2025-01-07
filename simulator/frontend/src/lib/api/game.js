import { env } from "$env/dynamic/public"

const API_URL = env["PUBLIC_API_URL"]

export async function get_water() {
    const url = `${API_URL}/game/water/current`;
    let response = await fetch(url, {
        method: "GET",
    });

    if (!response.ok) {
        throw new Error(`Response status: ${response.status} ${response.statusText}`);
    }

    /** @type {{statusCode: number, water: number}} */
    const json = await response.json();
    console.log(url);

    return json;
}

/**
 * Change the water level by `amount`.
 * @param {number} amount
 * @returns {Promise<void>}
 */
export async function change_water_by(amount) {
    const url = `${API_URL}/game/water/change-by?amount=${amount}`;
    let response = await fetch(url, {
        method: "POST",
    });
    console.log(response);
}