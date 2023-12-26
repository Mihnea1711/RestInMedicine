export function trimLastSegmentFromUrl(url) {
    // Remove the last part after the last '/'
    return url.slice(0, url.lastIndexOf('/') + 1);
}