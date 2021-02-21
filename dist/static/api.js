api = {
    _roomCache: [],
    getRooms: async () => {
        if (this._roomCache)
            return this._roomCache;
        let res = await (await fetch('/api/rooms')).json();
        if (!res.success)
            return res.reason;
        this._roomCache = res.rooms;
        return this._roomCache;
    },

}
