/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export class RpcStatus extends SpeakeasyBase {
    @SpeakeasyMetadata()
    @Expose({ name: "code" })
    code?: number;

    @SpeakeasyMetadata()
    @Expose({ name: "details" })
    details?: Record<string, any>[];

    @SpeakeasyMetadata()
    @Expose({ name: "message" })
    message?: string;
}
