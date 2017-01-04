import grpc
from grpc.framework.common import cardinality
from grpc.framework.interfaces.face import utilities as face_utilities

import ipersistence_pb2 as ipersistence__pb2
import ipersistence_pb2 as ipersistence__pb2
import ipersistence_pb2 as ipersistence__pb2
import ipersistence_pb2 as ipersistence__pb2


class SaverStub(object):
  """The greeting service definition.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.FetchSave = channel.unary_unary(
        '/ipersistence.Saver/FetchSave',
        request_serializer=ipersistence__pb2.FetchSaveRequest.SerializeToString,
        response_deserializer=ipersistence__pb2.FetchSaveReply.FromString,
        )
    self.UpdateTags = channel.unary_unary(
        '/ipersistence.Saver/UpdateTags',
        request_serializer=ipersistence__pb2.UpdateTagsRequest.SerializeToString,
        response_deserializer=ipersistence__pb2.BaseResultReply.FromString,
        )


class SaverServicer(object):
  """The greeting service definition.
  """

  def FetchSave(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateTags(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SaverServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'FetchSave': grpc.unary_unary_rpc_method_handler(
          servicer.FetchSave,
          request_deserializer=ipersistence__pb2.FetchSaveRequest.FromString,
          response_serializer=ipersistence__pb2.FetchSaveReply.SerializeToString,
      ),
      'UpdateTags': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateTags,
          request_deserializer=ipersistence__pb2.UpdateTagsRequest.FromString,
          response_serializer=ipersistence__pb2.BaseResultReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'ipersistence.Saver', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
