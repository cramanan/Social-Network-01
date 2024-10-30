export default function formatDate(timestamp: string) {
    return new Intl.DateTimeFormat("en-US", {
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
        day: "2-digit",
        month: "2-digit",
        year: "numeric",
        hour12: false, // Change to true for 12-hour format
    }).format(new Date(timestamp));
}
