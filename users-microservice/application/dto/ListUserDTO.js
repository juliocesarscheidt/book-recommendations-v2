class ListUserResponse {
  users;

  constructor(users) {
    const usersMapped = users.map(user => {
      const u = JSON.parse(JSON.stringify(user));
      u.uuid = u._id;
      delete u._id;
      if (u && u.created_at) u.created_at = this.mapDateToIsoString(u.created_at);
      if (u && u.updated_at) u.updated_at = this.mapDateToIsoString(u.updated_at);
      return u;
    });
    this.users = usersMapped;
  }

  mapDateToIsoString(date) {
    return date && date instanceof Date ? date.toISOString() : date;
  }
}

module.exports = ListUserResponse;
