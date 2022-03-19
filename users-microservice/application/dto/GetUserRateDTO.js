class GetUserRateDTO {
  userRate;

  constructor(userRate) {
    if (!userRate) return;

    const u = JSON.parse(JSON.stringify(userRate));
    u.user_uuid = u._id;
    delete u._id;
    if (u && u.created_at) u.created_at = this.mapDateToIsoString(u.created_at);
    if (u && u.updated_at) u.updated_at = this.mapDateToIsoString(u.updated_at);
    this.userRate = u;
  }

  mapDateToIsoString(date) {
    return date && date instanceof Date ? date.toISOString() : date;
  }
}

module.exports = GetUserRateDTO;
