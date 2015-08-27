# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vtctlservice.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import vtctldata_pb2 as vtctldata__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='vtctlservice.proto',
  package='vtctlservice',
  serialized_pb=_b('\n\x12vtctlservice.proto\x12\x0cvtctlservice\x1a\x0fvtctldata.proto2q\n\x05Vtctl\x12h\n\x13\x45xecuteVtctlCommand\x12%.vtctldata.ExecuteVtctlCommandRequest\x1a&.vtctldata.ExecuteVtctlCommandResponse\"\x00\x30\x01\x62\x06proto3')
  ,
  dependencies=[vtctldata__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)





import abc
from grpc.early_adopter import implementations
from grpc.framework.alpha import utilities
class EarlyAdopterVtctlServicer(object):
  """<fill me in later!>"""
  __metaclass__ = abc.ABCMeta
  @abc.abstractmethod
  def ExecuteVtctlCommand(self, request, context):
    raise NotImplementedError()
class EarlyAdopterVtctlServer(object):
  """<fill me in later!>"""
  __metaclass__ = abc.ABCMeta
  @abc.abstractmethod
  def start(self):
    raise NotImplementedError()
  @abc.abstractmethod
  def stop(self):
    raise NotImplementedError()
class EarlyAdopterVtctlStub(object):
  """<fill me in later!>"""
  __metaclass__ = abc.ABCMeta
  @abc.abstractmethod
  def ExecuteVtctlCommand(self, request):
    raise NotImplementedError()
  ExecuteVtctlCommand.async = None
def early_adopter_create_Vtctl_server(servicer, port, root_certificates, key_chain_pairs):
  import vtctldata_pb2
  import vtctldata_pb2
  method_service_descriptions = {
    "ExecuteVtctlCommand": utilities.unary_stream_service_description(
      servicer.ExecuteVtctlCommand,
      vtctldata_pb2.ExecuteVtctlCommandRequest.FromString,
      vtctldata_pb2.ExecuteVtctlCommandResponse.SerializeToString,
    ),
  }
  return implementations.secure_server("vtctlservice.Vtctl", method_service_descriptions, port, root_certificates, key_chain_pairs)
def early_adopter_create_Vtctl_stub(host, port):
  import vtctldata_pb2
  import vtctldata_pb2
  method_invocation_descriptions = {
    "ExecuteVtctlCommand": utilities.unary_stream_invocation_description(
      vtctldata_pb2.ExecuteVtctlCommandRequest.SerializeToString,
      vtctldata_pb2.ExecuteVtctlCommandResponse.FromString,
    ),
  }
  return implementations.insecure_stub("vtctlservice.Vtctl", method_invocation_descriptions, host, port)
# @@protoc_insertion_point(module_scope)
