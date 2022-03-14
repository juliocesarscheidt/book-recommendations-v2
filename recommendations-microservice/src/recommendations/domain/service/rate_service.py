# def build_rates(rates):
#     rates_dict = {}

#     for _rates in [rate["rates"] for rate in rates]:
#         for _rate in _rates:
#             rates_dict[_rate["user_uuid"]] = {}

#     for _rate in rates:
#         for rate_inner in _rate["rates"]:
#             rates_dict[rate_inner["user_uuid"]].update(
#                 {_rate["book_uuid"]: float(rate_inner["rate"])}
#             )

#     return rates_dict

def build_rates(rates):
    rates_dict = {}
    for rate in rates:
      inner_rates_dict = {}
      inner_rates = [{'book_uuid': rate['book_uuid'], 'rate': float(rate['rate'])} for rate in rate['rates']]
      for inner_rate in inner_rates:
        inner_rates_dict[inner_rate['book_uuid']] = inner_rate['rate']
      rates_dict[rate['user_uuid']] = inner_rates_dict
    return rates_dict
