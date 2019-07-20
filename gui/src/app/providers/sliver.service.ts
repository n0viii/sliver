/*
  Sliver Implant Framework
  Copyright (C) 2019  Bishop Fox
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.
--------------------------------------------------------------------------

This service is responsible for all of the Sliver Server interactions that
use protobuf.

*/

import { Injectable } from '@angular/core';
import { IPCService } from './ipc.service';
import { ProtobufService } from './protobuf.service';
import * as pb from '../../../rpc/pb';
import { GenerateReq } from '../../../rpc/pb';


@Injectable({
  providedIn: 'root'
})
export class SliverService extends ProtobufService {

  constructor(private _ipc: IPCService) {
    super();
  }

  sessions(): Promise<pb.Sessions> {
    return new Promise(async (resolve, reject) => {
      try {
        const reqEnvelope = new pb.Envelope();
        reqEnvelope.setType(pb.ClientPB.MsgSessions);
        const resp: string = await this._ipc.request('rpc_request', this.encode(reqEnvelope));
        resolve(pb.Sessions.deserializeBinary(this.decode(resp)));
      } catch (err) {
        reject(err);
      }
    });
  }

  sessionById(id: number): Promise<pb.Sliver> {
    return new Promise(async (resolve, reject) => {
      const sessions = await this.sessions();
      const slivers = sessions.getSliversList();
      for (let index = 0; index < slivers.length; ++index) {
        if (slivers[index].getId() === id) {
          resolve(slivers[index]);
          return;
        }
      }
      reject();
    });
  }

  sliverBuilds(): Promise<pb.SliverBuilds> {
    return new Promise(async (resolve) => {
      const reqEnvelope = new pb.Envelope();
      reqEnvelope.setType(pb.ClientPB.MsgListSliverBuilds);
      const resp: string = await this._ipc.request('rpc_request', this.encode(reqEnvelope));
      resolve(pb.SliverBuilds.deserializeBinary(this.decode(resp)));
    });
  }

  generate(config: pb.SliverConfig): Promise<pb.Generate> {
    return new Promise(async (resolve) => {
      const reqEnvelope = new pb.Envelope();
      const generateReq = new pb.GenerateReq();
      generateReq.setConfig(config);
      reqEnvelope.setType(pb.ClientPB.MsgGenerate);
      const resp: string = await this._ipc.request('rpc_request', this.encode(reqEnvelope));
      resolve(pb.Generate.deserializeBinary(this.decode(resp)));
    });
  }

  regenerate(name: string): Promise<pb.Regenerate> {
    return new Promise(async (resolve) => {
      const reqEnvelope = new pb.Envelope();
      reqEnvelope.setType(pb.ClientPB.MsgRegenerate);
      const regenReq = new pb.Regenerate();
      regenReq.setSlivername(name);
      const resp: string = await this._ipc.request('rpc_request', this.encode(reqEnvelope));
      resolve(pb.Regenerate.deserializeBinary(this.decode(resp)));
    });
  }

  ps(sliverId: number): Promise<pb.Ps> {
    return new Promise(async (resolve, reject) => {
      const reqEnvelope = new pb.Envelope();
      reqEnvelope.setType(pb.SliverPB.MsgPsReq);
      const psReq = new pb.PsReq();
      psReq.setSliverid(sliverId);
      reqEnvelope.setData(psReq.serializeBinary());
      const resp: string = await this._ipc.request('rpc_request', this.encode(reqEnvelope));
      resolve(pb.Ps.deserializeBinary(this.decode(resp)));
    });
  }

}
