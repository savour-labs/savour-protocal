# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: savourrpc/market.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from savourrpc import common_pb2 as savourrpc_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='savourrpc/market.proto',
  package='savourrpc.market',
  syntax='proto3',
  serialized_options=b'\n\024com.savourrpc.marketZ2git.savour.io/savour/savourrpc/go-savourrpc/market',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x16savourrpc/market.proto\x12\x10savourrpc.market\x1a\x16savourrpc/common.proto\"0\n\x06Symbol\x12\x11\n\tsymbol_id\x18\x01 \x01(\t\x12\x13\n\x0bsymbol_name\x18\x02 \x01(\t\"\'\n\rSymbolRequest\x12\x16\n\x0e\x63onsumer_token\x18\x01 \x01(\t\"m\n\x0eSymbolResponse\x12#\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x15.savourrpc.ReturnCode\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12)\n\x07symbols\x18\x03 \x03(\x0b\x32\x18.savourrpc.market.Symbol\"\x92\x01\n\x0bSymbolPrice\x12\x13\n\x0bsymbol_name\x18\x01 \x01(\t\x12\x11\n\tbuy_price\x18\x02 \x01(\t\x12\x12\n\nsell_price\x18\x03 \x01(\t\x12\x11\n\tavg_price\x18\x04 \x01(\t\x12\x11\n\tusd_price\x18\x05 \x01(\t\x12\x11\n\tcny_price\x18\x06 \x01(\t\x12\x0e\n\x06margin\x18\x07 \x01(\t\"?\n\x12SymbolPriceRequest\x12\x16\n\x0e\x63onsumer_token\x18\x01 \x01(\t\x12\x11\n\tsymbol_id\x18\x02 \x01(\t\"}\n\x13SymbolPriceResponse\x12#\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x15.savourrpc.ReturnCode\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x34\n\rsymbol_prices\x18\x03 \x03(\x0b\x32\x1d.savourrpc.market.SymbolPrice\"0\n\nStableCoin\x12\x0f\n\x07\x63oin_id\x18\x01 \x01(\t\x12\x11\n\tcoin_name\x18\x02 \x01(\t\"+\n\x11StableCoinRequest\x12\x16\n\x0e\x63onsumer_token\x18\x01 \x01(\t\"z\n\x12StableCoinResponse\x12#\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x15.savourrpc.ReturnCode\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x32\n\x0cstable_coins\x18\x03 \x03(\x0b\x32\x1c.savourrpc.market.StableCoin\"Q\n\x0fStableCoinPrice\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x11\n\tusd_price\x18\x03 \x01(\t\x12\x11\n\tcny_price\x18\x04 \x01(\t\"A\n\x16StableCoinPriceRequest\x12\x16\n\x0e\x63onsumer_token\x18\x01 \x01(\t\x12\x0f\n\x07\x63oin_id\x18\x02 \x01(\t\"\x83\x01\n\x17StableCoinPriceResponse\x12#\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x15.savourrpc.ReturnCode\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x36\n\x0b\x63oin_prices\x18\x03 \x03(\x0b\x32!.savourrpc.market.StableCoinPrice2\x8f\x03\n\x0cPriceService\x12Q\n\ngetSymbols\x12\x1f.savourrpc.market.SymbolRequest\x1a .savourrpc.market.SymbolResponse\"\x00\x12`\n\x0fgetSymbolPrices\x12$.savourrpc.market.SymbolPriceRequest\x1a%.savourrpc.market.SymbolPriceResponse\"\x00\x12]\n\x0egetStableCoins\x12#.savourrpc.market.StableCoinRequest\x1a$.savourrpc.market.StableCoinResponse\"\x00\x12k\n\x12getStableCoinPrice\x12(.savourrpc.market.StableCoinPriceRequest\x1a).savourrpc.market.StableCoinPriceResponse\"\x00\x42J\n\x14\x63om.savourrpc.marketZ2git.savour.io/savour/savourrpc/go-savourrpc/marketb\x06proto3'
  ,
  dependencies=[savourrpc_dot_common__pb2.DESCRIPTOR,])




_SYMBOL = _descriptor.Descriptor(
  name='Symbol',
  full_name='savourrpc.market.Symbol',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='symbol_id', full_name='savourrpc.market.Symbol.symbol_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='symbol_name', full_name='savourrpc.market.Symbol.symbol_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=68,
  serialized_end=116,
)


_SYMBOLREQUEST = _descriptor.Descriptor(
  name='SymbolRequest',
  full_name='savourrpc.market.SymbolRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='consumer_token', full_name='savourrpc.market.SymbolRequest.consumer_token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=118,
  serialized_end=157,
)


_SYMBOLRESPONSE = _descriptor.Descriptor(
  name='SymbolResponse',
  full_name='savourrpc.market.SymbolResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='savourrpc.market.SymbolResponse.code', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='savourrpc.market.SymbolResponse.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='symbols', full_name='savourrpc.market.SymbolResponse.symbols', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=159,
  serialized_end=268,
)


_SYMBOLPRICE = _descriptor.Descriptor(
  name='SymbolPrice',
  full_name='savourrpc.market.SymbolPrice',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='symbol_name', full_name='savourrpc.market.SymbolPrice.symbol_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='buy_price', full_name='savourrpc.market.SymbolPrice.buy_price', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sell_price', full_name='savourrpc.market.SymbolPrice.sell_price', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='avg_price', full_name='savourrpc.market.SymbolPrice.avg_price', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='usd_price', full_name='savourrpc.market.SymbolPrice.usd_price', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cny_price', full_name='savourrpc.market.SymbolPrice.cny_price', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='margin', full_name='savourrpc.market.SymbolPrice.margin', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=271,
  serialized_end=417,
)


_SYMBOLPRICEREQUEST = _descriptor.Descriptor(
  name='SymbolPriceRequest',
  full_name='savourrpc.market.SymbolPriceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='consumer_token', full_name='savourrpc.market.SymbolPriceRequest.consumer_token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='symbol_id', full_name='savourrpc.market.SymbolPriceRequest.symbol_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=419,
  serialized_end=482,
)


_SYMBOLPRICERESPONSE = _descriptor.Descriptor(
  name='SymbolPriceResponse',
  full_name='savourrpc.market.SymbolPriceResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='savourrpc.market.SymbolPriceResponse.code', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='savourrpc.market.SymbolPriceResponse.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='symbol_prices', full_name='savourrpc.market.SymbolPriceResponse.symbol_prices', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=484,
  serialized_end=609,
)


_STABLECOIN = _descriptor.Descriptor(
  name='StableCoin',
  full_name='savourrpc.market.StableCoin',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='coin_id', full_name='savourrpc.market.StableCoin.coin_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='coin_name', full_name='savourrpc.market.StableCoin.coin_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=611,
  serialized_end=659,
)


_STABLECOINREQUEST = _descriptor.Descriptor(
  name='StableCoinRequest',
  full_name='savourrpc.market.StableCoinRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='consumer_token', full_name='savourrpc.market.StableCoinRequest.consumer_token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=661,
  serialized_end=704,
)


_STABLECOINRESPONSE = _descriptor.Descriptor(
  name='StableCoinResponse',
  full_name='savourrpc.market.StableCoinResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='savourrpc.market.StableCoinResponse.code', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='savourrpc.market.StableCoinResponse.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='stable_coins', full_name='savourrpc.market.StableCoinResponse.stable_coins', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=706,
  serialized_end=828,
)


_STABLECOINPRICE = _descriptor.Descriptor(
  name='StableCoinPrice',
  full_name='savourrpc.market.StableCoinPrice',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='savourrpc.market.StableCoinPrice.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='savourrpc.market.StableCoinPrice.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='usd_price', full_name='savourrpc.market.StableCoinPrice.usd_price', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cny_price', full_name='savourrpc.market.StableCoinPrice.cny_price', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=830,
  serialized_end=911,
)


_STABLECOINPRICEREQUEST = _descriptor.Descriptor(
  name='StableCoinPriceRequest',
  full_name='savourrpc.market.StableCoinPriceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='consumer_token', full_name='savourrpc.market.StableCoinPriceRequest.consumer_token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='coin_id', full_name='savourrpc.market.StableCoinPriceRequest.coin_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=913,
  serialized_end=978,
)


_STABLECOINPRICERESPONSE = _descriptor.Descriptor(
  name='StableCoinPriceResponse',
  full_name='savourrpc.market.StableCoinPriceResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='savourrpc.market.StableCoinPriceResponse.code', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='savourrpc.market.StableCoinPriceResponse.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='coin_prices', full_name='savourrpc.market.StableCoinPriceResponse.coin_prices', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=981,
  serialized_end=1112,
)

_SYMBOLRESPONSE.fields_by_name['code'].enum_type = savourrpc_dot_common__pb2._RETURNCODE
_SYMBOLRESPONSE.fields_by_name['symbols'].message_type = _SYMBOL
_SYMBOLPRICERESPONSE.fields_by_name['code'].enum_type = savourrpc_dot_common__pb2._RETURNCODE
_SYMBOLPRICERESPONSE.fields_by_name['symbol_prices'].message_type = _SYMBOLPRICE
_STABLECOINRESPONSE.fields_by_name['code'].enum_type = savourrpc_dot_common__pb2._RETURNCODE
_STABLECOINRESPONSE.fields_by_name['stable_coins'].message_type = _STABLECOIN
_STABLECOINPRICERESPONSE.fields_by_name['code'].enum_type = savourrpc_dot_common__pb2._RETURNCODE
_STABLECOINPRICERESPONSE.fields_by_name['coin_prices'].message_type = _STABLECOINPRICE
DESCRIPTOR.message_types_by_name['Symbol'] = _SYMBOL
DESCRIPTOR.message_types_by_name['SymbolRequest'] = _SYMBOLREQUEST
DESCRIPTOR.message_types_by_name['SymbolResponse'] = _SYMBOLRESPONSE
DESCRIPTOR.message_types_by_name['SymbolPrice'] = _SYMBOLPRICE
DESCRIPTOR.message_types_by_name['SymbolPriceRequest'] = _SYMBOLPRICEREQUEST
DESCRIPTOR.message_types_by_name['SymbolPriceResponse'] = _SYMBOLPRICERESPONSE
DESCRIPTOR.message_types_by_name['StableCoin'] = _STABLECOIN
DESCRIPTOR.message_types_by_name['StableCoinRequest'] = _STABLECOINREQUEST
DESCRIPTOR.message_types_by_name['StableCoinResponse'] = _STABLECOINRESPONSE
DESCRIPTOR.message_types_by_name['StableCoinPrice'] = _STABLECOINPRICE
DESCRIPTOR.message_types_by_name['StableCoinPriceRequest'] = _STABLECOINPRICEREQUEST
DESCRIPTOR.message_types_by_name['StableCoinPriceResponse'] = _STABLECOINPRICERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Symbol = _reflection.GeneratedProtocolMessageType('Symbol', (_message.Message,), {
  'DESCRIPTOR' : _SYMBOL,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.Symbol)
  })
_sym_db.RegisterMessage(Symbol)

SymbolRequest = _reflection.GeneratedProtocolMessageType('SymbolRequest', (_message.Message,), {
  'DESCRIPTOR' : _SYMBOLREQUEST,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.SymbolRequest)
  })
_sym_db.RegisterMessage(SymbolRequest)

SymbolResponse = _reflection.GeneratedProtocolMessageType('SymbolResponse', (_message.Message,), {
  'DESCRIPTOR' : _SYMBOLRESPONSE,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.SymbolResponse)
  })
_sym_db.RegisterMessage(SymbolResponse)

SymbolPrice = _reflection.GeneratedProtocolMessageType('SymbolPrice', (_message.Message,), {
  'DESCRIPTOR' : _SYMBOLPRICE,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.SymbolPrice)
  })
_sym_db.RegisterMessage(SymbolPrice)

SymbolPriceRequest = _reflection.GeneratedProtocolMessageType('SymbolPriceRequest', (_message.Message,), {
  'DESCRIPTOR' : _SYMBOLPRICEREQUEST,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.SymbolPriceRequest)
  })
_sym_db.RegisterMessage(SymbolPriceRequest)

SymbolPriceResponse = _reflection.GeneratedProtocolMessageType('SymbolPriceResponse', (_message.Message,), {
  'DESCRIPTOR' : _SYMBOLPRICERESPONSE,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.SymbolPriceResponse)
  })
_sym_db.RegisterMessage(SymbolPriceResponse)

StableCoin = _reflection.GeneratedProtocolMessageType('StableCoin', (_message.Message,), {
  'DESCRIPTOR' : _STABLECOIN,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.StableCoin)
  })
_sym_db.RegisterMessage(StableCoin)

StableCoinRequest = _reflection.GeneratedProtocolMessageType('StableCoinRequest', (_message.Message,), {
  'DESCRIPTOR' : _STABLECOINREQUEST,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.StableCoinRequest)
  })
_sym_db.RegisterMessage(StableCoinRequest)

StableCoinResponse = _reflection.GeneratedProtocolMessageType('StableCoinResponse', (_message.Message,), {
  'DESCRIPTOR' : _STABLECOINRESPONSE,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.StableCoinResponse)
  })
_sym_db.RegisterMessage(StableCoinResponse)

StableCoinPrice = _reflection.GeneratedProtocolMessageType('StableCoinPrice', (_message.Message,), {
  'DESCRIPTOR' : _STABLECOINPRICE,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.StableCoinPrice)
  })
_sym_db.RegisterMessage(StableCoinPrice)

StableCoinPriceRequest = _reflection.GeneratedProtocolMessageType('StableCoinPriceRequest', (_message.Message,), {
  'DESCRIPTOR' : _STABLECOINPRICEREQUEST,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.StableCoinPriceRequest)
  })
_sym_db.RegisterMessage(StableCoinPriceRequest)

StableCoinPriceResponse = _reflection.GeneratedProtocolMessageType('StableCoinPriceResponse', (_message.Message,), {
  'DESCRIPTOR' : _STABLECOINPRICERESPONSE,
  '__module__' : 'savourrpc.market_pb2'
  # @@protoc_insertion_point(class_scope:savourrpc.market.StableCoinPriceResponse)
  })
_sym_db.RegisterMessage(StableCoinPriceResponse)


DESCRIPTOR._options = None

_PRICESERVICE = _descriptor.ServiceDescriptor(
  name='PriceService',
  full_name='savourrpc.market.PriceService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1115,
  serialized_end=1514,
  methods=[
  _descriptor.MethodDescriptor(
    name='getSymbols',
    full_name='savourrpc.market.PriceService.getSymbols',
    index=0,
    containing_service=None,
    input_type=_SYMBOLREQUEST,
    output_type=_SYMBOLRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='getSymbolPrices',
    full_name='savourrpc.market.PriceService.getSymbolPrices',
    index=1,
    containing_service=None,
    input_type=_SYMBOLPRICEREQUEST,
    output_type=_SYMBOLPRICERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='getStableCoins',
    full_name='savourrpc.market.PriceService.getStableCoins',
    index=2,
    containing_service=None,
    input_type=_STABLECOINREQUEST,
    output_type=_STABLECOINRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='getStableCoinPrice',
    full_name='savourrpc.market.PriceService.getStableCoinPrice',
    index=3,
    containing_service=None,
    input_type=_STABLECOINPRICEREQUEST,
    output_type=_STABLECOINPRICERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_PRICESERVICE)

DESCRIPTOR.services_by_name['PriceService'] = _PRICESERVICE

# @@protoc_insertion_point(module_scope)
