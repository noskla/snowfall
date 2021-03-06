api = {
    _roomCache: [],
    getRooms: async () => {
        if (!this._roomCache) {
            let res = await (await fetch('/api/rooms')).json();
            if (!res.success)
                return res.reason;
            this._roomCache = res.rooms;
        }
        return this._roomCache;
    },

    _standCache: [],
    getStands: async () => {
        if (!this._standCache) {
            let res = await (await fetch('/api/stands')).json();
            if (!res.success)
                return res.reason;
            this._standCache = res.stands;
        }
        return this._standCache
    },

    newUser: async (username, password, discord) => {
        return await (await fetch('/api/user', {
            method: 'POST', headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                username: username, password: password, discord: discord})})).json();
    },

    confirmDiscord: async (userID, confirmCode) => {
        return await (await fetch(`/api/user/${userID}/discord`, {
            method: 'POST', headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({discordKey: confirmCode})})).json();
    }
}
