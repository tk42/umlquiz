import { API_URL, ENDPOINT_TOKEN } from './const';

export const getToken = () => {
    return fetch(API_URL+ENDPOINT_TOKEN, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: "username="+process.env.AUTH_USERNAME+"&password="+process.env.AUTH_PASSWORD,
    }).then((data) => {
        return data.json()
    }).then((map: Map<string, string>) => {
        return map.get("access_token");
    }).then((token?: string) => {
        return token===undefined?"":token
    });
}
