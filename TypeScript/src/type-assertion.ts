interface APIErr extends Error {
    code: number,
}


interface HTTPErr extends Error {
    statusCode: number,
}


function isAPIError(err: Error) {
    if (err as APIErr) {
        return true;
    }

    return false;
}
