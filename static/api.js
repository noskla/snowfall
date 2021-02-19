api = {
    getRooms: async () => {
        let res = await fetch('/api/rooms');
        return await res.json();
    },
}
