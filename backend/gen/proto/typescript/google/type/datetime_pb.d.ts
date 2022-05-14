// package: google.type
// file: google/type/datetime.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";

export class DateTime extends jspb.Message { 
    getYear(): number;
    setYear(value: number): DateTime;
    getMonth(): number;
    setMonth(value: number): DateTime;
    getDay(): number;
    setDay(value: number): DateTime;
    getHours(): number;
    setHours(value: number): DateTime;
    getMinutes(): number;
    setMinutes(value: number): DateTime;
    getSeconds(): number;
    setSeconds(value: number): DateTime;
    getNanos(): number;
    setNanos(value: number): DateTime;

    hasUtcOffset(): boolean;
    clearUtcOffset(): void;
    getUtcOffset(): google_protobuf_duration_pb.Duration | undefined;
    setUtcOffset(value?: google_protobuf_duration_pb.Duration): DateTime;

    hasTimeZone(): boolean;
    clearTimeZone(): void;
    getTimeZone(): TimeZone | undefined;
    setTimeZone(value?: TimeZone): DateTime;

    getTimeOffsetCase(): DateTime.TimeOffsetCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DateTime.AsObject;
    static toObject(includeInstance: boolean, msg: DateTime): DateTime.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DateTime, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DateTime;
    static deserializeBinaryFromReader(message: DateTime, reader: jspb.BinaryReader): DateTime;
}

export namespace DateTime {
    export type AsObject = {
        year: number,
        month: number,
        day: number,
        hours: number,
        minutes: number,
        seconds: number,
        nanos: number,
        utcOffset?: google_protobuf_duration_pb.Duration.AsObject,
        timeZone?: TimeZone.AsObject,
    }

    export enum TimeOffsetCase {
        TIME_OFFSET_NOT_SET = 0,
        UTC_OFFSET = 8,
        TIME_ZONE = 9,
    }

}

export class TimeZone extends jspb.Message { 
    getId(): string;
    setId(value: string): TimeZone;
    getVersion(): string;
    setVersion(value: string): TimeZone;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TimeZone.AsObject;
    static toObject(includeInstance: boolean, msg: TimeZone): TimeZone.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TimeZone, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TimeZone;
    static deserializeBinaryFromReader(message: TimeZone, reader: jspb.BinaryReader): TimeZone;
}

export namespace TimeZone {
    export type AsObject = {
        id: string,
        version: string,
    }
}
