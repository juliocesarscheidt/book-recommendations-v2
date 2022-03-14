class UpsertUserRateDTO {
  user_uuid;
  constructor(_id) {
    this.user_uuid = _id;
  }
}

module.exports = UpsertUserRateDTO;
