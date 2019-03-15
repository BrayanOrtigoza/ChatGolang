export async function postService(route, body, headers) {
    try {
        let response = await fetch(
            route, {
                method: 'POST',
                headers: headers,
                body: JSON.stringify(body)
            }
        );
        let responseJson = await response.json();
        return responseJson;
    } catch (error) {
        alert(error);
    }
}

export async function getService(route, headers) {
    try {
        let response = await fetch(
            route, {
                method: 'GET',
                headers: headers
            }
        );
        let responseJson = await response.json();
        return responseJson;
    } catch (error) {
        alert(error);
    }
}