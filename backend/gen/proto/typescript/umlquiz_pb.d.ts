// package: umlquiz
// file: umlquiz.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_type_datetime_pb from "./google/type/datetime_pb";

export class User extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): User;
    getUsername(): string;
    setUsername(value: string): User;
    getPassword(): string;
    setPassword(value: string): User;
    getEmail(): string;
    setEmail(value: string): User;
    getProfile(): string;
    setProfile(value: string): User;
    getMembership(): MemgerShip;
    setMembership(value: MemgerShip): User;
    clearLikedQuizIdsList(): void;
    getLikedQuizIdsList(): Array<string>;
    setLikedQuizIdsList(value: Array<string>): User;
    addLikedQuizIds(value: string, index?: number): string;
    clearQuizHistoryList(): void;
    getQuizHistoryList(): Array<string>;
    setQuizHistoryList(value: Array<string>): User;
    addQuizHistory(value: string, index?: number): string;

    hasCreatedAt(): boolean;
    clearCreatedAt(): void;
    getCreatedAt(): google_type_datetime_pb.DateTime | undefined;
    setCreatedAt(value?: google_type_datetime_pb.DateTime): User;

    hasUpdatedAt(): boolean;
    clearUpdatedAt(): void;
    getUpdatedAt(): google_type_datetime_pb.DateTime | undefined;
    setUpdatedAt(value?: google_type_datetime_pb.DateTime): User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): User.AsObject;
    static toObject(includeInstance: boolean, msg: User): User.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): User;
    static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
    export type AsObject = {
        userId: string,
        username: string,
        password: string,
        email: string,
        profile: string,
        membership: MemgerShip,
        likedQuizIdsList: Array<string>,
        quizHistoryList: Array<string>,
        createdAt?: google_type_datetime_pb.DateTime.AsObject,
        updatedAt?: google_type_datetime_pb.DateTime.AsObject,
    }
}

export class Quiz extends jspb.Message { 
    getQuizId(): string;
    setQuizId(value: string): Quiz;
    getLanguage(): string;
    setLanguage(value: string): Quiz;
    getStatus(): QuizStatus;
    setStatus(value: QuizStatus): Quiz;
    getDiagramType(): DiagramType;
    setDiagramType(value: DiagramType): Quiz;
    getLevel(): string;
    setLevel(value: string): Quiz;
    getTitle(): string;
    setTitle(value: string): Quiz;
    getText(): string;
    setText(value: string): Quiz;
    getDiagram(): string;
    setDiagram(value: string): Quiz;
    getLikes(): number;
    setLikes(value: number): Quiz;
    getAuthorId(): string;
    setAuthorId(value: string): Quiz;

    hasCreatedAt(): boolean;
    clearCreatedAt(): void;
    getCreatedAt(): google_type_datetime_pb.DateTime | undefined;
    setCreatedAt(value?: google_type_datetime_pb.DateTime): Quiz;

    hasUpdatedAt(): boolean;
    clearUpdatedAt(): void;
    getUpdatedAt(): google_type_datetime_pb.DateTime | undefined;
    setUpdatedAt(value?: google_type_datetime_pb.DateTime): Quiz;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Quiz.AsObject;
    static toObject(includeInstance: boolean, msg: Quiz): Quiz.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Quiz, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Quiz;
    static deserializeBinaryFromReader(message: Quiz, reader: jspb.BinaryReader): Quiz;
}

export namespace Quiz {
    export type AsObject = {
        quizId: string,
        language: string,
        status: QuizStatus,
        diagramType: DiagramType,
        level: string,
        title: string,
        text: string,
        diagram: string,
        likes: number,
        authorId: string,
        createdAt?: google_type_datetime_pb.DateTime.AsObject,
        updatedAt?: google_type_datetime_pb.DateTime.AsObject,
    }
}

export class Report extends jspb.Message { 
    getReportId(): string;
    setReportId(value: string): Report;
    getQuizId(): string;
    setQuizId(value: string): Report;
    getLanguage(): string;
    setLanguage(value: string): Report;
    getAuthorId(): string;
    setAuthorId(value: string): Report;
    getTitle(): string;
    setTitle(value: string): Report;
    getText(): string;
    setText(value: string): Report;
    getDiagram(): string;
    setDiagram(value: string): Report;
    getComment(): string;
    setComment(value: string): Report;

    hasCreatedAt(): boolean;
    clearCreatedAt(): void;
    getCreatedAt(): google_type_datetime_pb.DateTime | undefined;
    setCreatedAt(value?: google_type_datetime_pb.DateTime): Report;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Report.AsObject;
    static toObject(includeInstance: boolean, msg: Report): Report.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Report, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Report;
    static deserializeBinaryFromReader(message: Report, reader: jspb.BinaryReader): Report;
}

export namespace Report {
    export type AsObject = {
        reportId: string,
        quizId: string,
        language: string,
        authorId: string,
        title: string,
        text: string,
        diagram: string,
        comment: string,
        createdAt?: google_type_datetime_pb.DateTime.AsObject,
    }
}

export class AddUserRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): AddUserRequest;
    getPassword(): string;
    setPassword(value: string): AddUserRequest;
    getEmail(): string;
    setEmail(value: string): AddUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AddUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AddUserRequest): AddUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AddUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AddUserRequest;
    static deserializeBinaryFromReader(message: AddUserRequest, reader: jspb.BinaryReader): AddUserRequest;
}

export namespace AddUserRequest {
    export type AsObject = {
        username: string,
        password: string,
        email: string,
    }
}

export class UpdateUserRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): UpdateUserRequest;
    getUsername(): string;
    setUsername(value: string): UpdateUserRequest;
    getPassword(): string;
    setPassword(value: string): UpdateUserRequest;
    getEmail(): string;
    setEmail(value: string): UpdateUserRequest;
    getProfile(): string;
    setProfile(value: string): UpdateUserRequest;
    getMembership(): MemgerShip;
    setMembership(value: MemgerShip): UpdateUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
    static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
    export type AsObject = {
        userId: string,
        username: string,
        password: string,
        email: string,
        profile: string,
        membership: MemgerShip,
    }
}

export class UserRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): UserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UserRequest): UserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserRequest;
    static deserializeBinaryFromReader(message: UserRequest, reader: jspb.BinaryReader): UserRequest;
}

export namespace UserRequest {
    export type AsObject = {
        userId: string,
    }
}

export class AddQuizRequest extends jspb.Message { 
    getLanguage(): string;
    setLanguage(value: string): AddQuizRequest;
    getDiagramType(): DiagramType;
    setDiagramType(value: DiagramType): AddQuizRequest;
    getLevel(): string;
    setLevel(value: string): AddQuizRequest;
    getTitle(): string;
    setTitle(value: string): AddQuizRequest;
    getText(): string;
    setText(value: string): AddQuizRequest;
    getDiagram(): string;
    setDiagram(value: string): AddQuizRequest;
    getAuthorId(): string;
    setAuthorId(value: string): AddQuizRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AddQuizRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AddQuizRequest): AddQuizRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AddQuizRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AddQuizRequest;
    static deserializeBinaryFromReader(message: AddQuizRequest, reader: jspb.BinaryReader): AddQuizRequest;
}

export namespace AddQuizRequest {
    export type AsObject = {
        language: string,
        diagramType: DiagramType,
        level: string,
        title: string,
        text: string,
        diagram: string,
        authorId: string,
    }
}

export class FindQuizRequest extends jspb.Message { 
    getQuizId(): string;
    setQuizId(value: string): FindQuizRequest;
    getLanguage(): string;
    setLanguage(value: string): FindQuizRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindQuizRequest.AsObject;
    static toObject(includeInstance: boolean, msg: FindQuizRequest): FindQuizRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindQuizRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindQuizRequest;
    static deserializeBinaryFromReader(message: FindQuizRequest, reader: jspb.BinaryReader): FindQuizRequest;
}

export namespace FindQuizRequest {
    export type AsObject = {
        quizId: string,
        language: string,
    }
}

export class UpdateQuizRequest extends jspb.Message { 
    getQuizId(): string;
    setQuizId(value: string): UpdateQuizRequest;
    getLanguage(): string;
    setLanguage(value: string): UpdateQuizRequest;
    getStatus(): QuizStatus;
    setStatus(value: QuizStatus): UpdateQuizRequest;
    getDiagramType(): DiagramType;
    setDiagramType(value: DiagramType): UpdateQuizRequest;
    getLevel(): string;
    setLevel(value: string): UpdateQuizRequest;
    getTitle(): string;
    setTitle(value: string): UpdateQuizRequest;
    getText(): string;
    setText(value: string): UpdateQuizRequest;
    getDiagram(): string;
    setDiagram(value: string): UpdateQuizRequest;
    getLikes(): number;
    setLikes(value: number): UpdateQuizRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateQuizRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateQuizRequest): UpdateQuizRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateQuizRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateQuizRequest;
    static deserializeBinaryFromReader(message: UpdateQuizRequest, reader: jspb.BinaryReader): UpdateQuizRequest;
}

export namespace UpdateQuizRequest {
    export type AsObject = {
        quizId: string,
        language: string,
        status: QuizStatus,
        diagramType: DiagramType,
        level: string,
        title: string,
        text: string,
        diagram: string,
        likes: number,
    }
}

export class DeleteQuizRequest extends jspb.Message { 
    getQuizId(): string;
    setQuizId(value: string): DeleteQuizRequest;
    getLanguage(): string;
    setLanguage(value: string): DeleteQuizRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteQuizRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteQuizRequest): DeleteQuizRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteQuizRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteQuizRequest;
    static deserializeBinaryFromReader(message: DeleteQuizRequest, reader: jspb.BinaryReader): DeleteQuizRequest;
}

export namespace DeleteQuizRequest {
    export type AsObject = {
        quizId: string,
        language: string,
    }
}

export class ListQuizzesAllRequest extends jspb.Message { 
    getLanguage(): string;
    setLanguage(value: string): ListQuizzesAllRequest;
    getStatus(): QuizStatus;
    setStatus(value: QuizStatus): ListQuizzesAllRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListQuizzesAllRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListQuizzesAllRequest): ListQuizzesAllRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListQuizzesAllRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListQuizzesAllRequest;
    static deserializeBinaryFromReader(message: ListQuizzesAllRequest, reader: jspb.BinaryReader): ListQuizzesAllRequest;
}

export namespace ListQuizzesAllRequest {
    export type AsObject = {
        language: string,
        status: QuizStatus,
    }
}

export class ListQuizzesByUserRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ListQuizzesByUserRequest;
    getLanguage(): string;
    setLanguage(value: string): ListQuizzesByUserRequest;
    getStatus(): QuizStatus;
    setStatus(value: QuizStatus): ListQuizzesByUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListQuizzesByUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListQuizzesByUserRequest): ListQuizzesByUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListQuizzesByUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListQuizzesByUserRequest;
    static deserializeBinaryFromReader(message: ListQuizzesByUserRequest, reader: jspb.BinaryReader): ListQuizzesByUserRequest;
}

export namespace ListQuizzesByUserRequest {
    export type AsObject = {
        userId: string,
        language: string,
        status: QuizStatus,
    }
}

export class SolveQuizRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): SolveQuizRequest;
    getQuizId(): string;
    setQuizId(value: string): SolveQuizRequest;
    getLanguage(): string;
    setLanguage(value: string): SolveQuizRequest;
    getDiagram(): string;
    setDiagram(value: string): SolveQuizRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SolveQuizRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SolveQuizRequest): SolveQuizRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SolveQuizRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SolveQuizRequest;
    static deserializeBinaryFromReader(message: SolveQuizRequest, reader: jspb.BinaryReader): SolveQuizRequest;
}

export namespace SolveQuizRequest {
    export type AsObject = {
        userId: string,
        quizId: string,
        language: string,
        diagram: string,
    }
}

export class LikeQuizRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): LikeQuizRequest;
    getQuizId(): string;
    setQuizId(value: string): LikeQuizRequest;
    getLanguage(): string;
    setLanguage(value: string): LikeQuizRequest;
    getDiffLike(): number;
    setDiffLike(value: number): LikeQuizRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LikeQuizRequest.AsObject;
    static toObject(includeInstance: boolean, msg: LikeQuizRequest): LikeQuizRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LikeQuizRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LikeQuizRequest;
    static deserializeBinaryFromReader(message: LikeQuizRequest, reader: jspb.BinaryReader): LikeQuizRequest;
}

export namespace LikeQuizRequest {
    export type AsObject = {
        userId: string,
        quizId: string,
        language: string,
        diffLike: number,
    }
}

export class AddReportRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): AddReportRequest;
    getQuizId(): string;
    setQuizId(value: string): AddReportRequest;
    getLanguage(): string;
    setLanguage(value: string): AddReportRequest;
    getTitle(): string;
    setTitle(value: string): AddReportRequest;
    getText(): string;
    setText(value: string): AddReportRequest;
    getDiagram(): string;
    setDiagram(value: string): AddReportRequest;
    getComment(): string;
    setComment(value: string): AddReportRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AddReportRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AddReportRequest): AddReportRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AddReportRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AddReportRequest;
    static deserializeBinaryFromReader(message: AddReportRequest, reader: jspb.BinaryReader): AddReportRequest;
}

export namespace AddReportRequest {
    export type AsObject = {
        userId: string,
        quizId: string,
        language: string,
        title: string,
        text: string,
        diagram: string,
        comment: string,
    }
}

export class FindReportsRequest extends jspb.Message { 
    getQuizId(): string;
    setQuizId(value: string): FindReportsRequest;
    getLanguage(): string;
    setLanguage(value: string): FindReportsRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FindReportsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: FindReportsRequest): FindReportsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FindReportsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FindReportsRequest;
    static deserializeBinaryFromReader(message: FindReportsRequest, reader: jspb.BinaryReader): FindReportsRequest;
}

export namespace FindReportsRequest {
    export type AsObject = {
        quizId: string,
        language: string,
    }
}

export class UpdateReportRequest extends jspb.Message { 
    getReportId(): string;
    setReportId(value: string): UpdateReportRequest;
    getTitle(): string;
    setTitle(value: string): UpdateReportRequest;
    getText(): string;
    setText(value: string): UpdateReportRequest;
    getDiagram(): string;
    setDiagram(value: string): UpdateReportRequest;
    getComment(): string;
    setComment(value: string): UpdateReportRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateReportRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateReportRequest): UpdateReportRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateReportRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateReportRequest;
    static deserializeBinaryFromReader(message: UpdateReportRequest, reader: jspb.BinaryReader): UpdateReportRequest;
}

export namespace UpdateReportRequest {
    export type AsObject = {
        reportId: string,
        title: string,
        text: string,
        diagram: string,
        comment: string,
    }
}

export class DeleteReportRequest extends jspb.Message { 
    getReportId(): string;
    setReportId(value: string): DeleteReportRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteReportRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteReportRequest): DeleteReportRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteReportRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteReportRequest;
    static deserializeBinaryFromReader(message: DeleteReportRequest, reader: jspb.BinaryReader): DeleteReportRequest;
}

export namespace DeleteReportRequest {
    export type AsObject = {
        reportId: string,
    }
}

export class UserResponse extends jspb.Message { 

    hasUser(): boolean;
    clearUser(): void;
    getUser(): User | undefined;
    setUser(value?: User): UserResponse;
    getError(): string;
    setError(value: string): UserResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UserResponse): UserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserResponse;
    static deserializeBinaryFromReader(message: UserResponse, reader: jspb.BinaryReader): UserResponse;
}

export namespace UserResponse {
    export type AsObject = {
        user?: User.AsObject,
        error: string,
    }
}

export class QuizResponse extends jspb.Message { 

    hasQuiz(): boolean;
    clearQuiz(): void;
    getQuiz(): Quiz | undefined;
    setQuiz(value?: Quiz): QuizResponse;
    getError(): string;
    setError(value: string): QuizResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QuizResponse.AsObject;
    static toObject(includeInstance: boolean, msg: QuizResponse): QuizResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QuizResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QuizResponse;
    static deserializeBinaryFromReader(message: QuizResponse, reader: jspb.BinaryReader): QuizResponse;
}

export namespace QuizResponse {
    export type AsObject = {
        quiz?: Quiz.AsObject,
        error: string,
    }
}

export class ReportResponse extends jspb.Message { 

    hasReport(): boolean;
    clearReport(): void;
    getReport(): Report | undefined;
    setReport(value?: Report): ReportResponse;
    getError(): string;
    setError(value: string): ReportResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReportResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ReportResponse): ReportResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReportResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReportResponse;
    static deserializeBinaryFromReader(message: ReportResponse, reader: jspb.BinaryReader): ReportResponse;
}

export namespace ReportResponse {
    export type AsObject = {
        report?: Report.AsObject,
        error: string,
    }
}

export class QuizzesResponse extends jspb.Message { 
    clearQuizList(): void;
    getQuizList(): Array<Quiz>;
    setQuizList(value: Array<Quiz>): QuizzesResponse;
    addQuiz(value?: Quiz, index?: number): Quiz;
    getError(): string;
    setError(value: string): QuizzesResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QuizzesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: QuizzesResponse): QuizzesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QuizzesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QuizzesResponse;
    static deserializeBinaryFromReader(message: QuizzesResponse, reader: jspb.BinaryReader): QuizzesResponse;
}

export namespace QuizzesResponse {
    export type AsObject = {
        quizList: Array<Quiz.AsObject>,
        error: string,
    }
}

export class ReportsResponse extends jspb.Message { 
    clearReportList(): void;
    getReportList(): Array<Report>;
    setReportList(value: Array<Report>): ReportsResponse;
    addReport(value?: Report, index?: number): Report;
    getError(): string;
    setError(value: string): ReportsResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReportsResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ReportsResponse): ReportsResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReportsResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReportsResponse;
    static deserializeBinaryFromReader(message: ReportsResponse, reader: jspb.BinaryReader): ReportsResponse;
}

export namespace ReportsResponse {
    export type AsObject = {
        reportList: Array<Report.AsObject>,
        error: string,
    }
}

export class SolveResponse extends jspb.Message { 
    getDiff(): string;
    setDiff(value: string): SolveResponse;
    getError(): string;
    setError(value: string): SolveResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SolveResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SolveResponse): SolveResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SolveResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SolveResponse;
    static deserializeBinaryFromReader(message: SolveResponse, reader: jspb.BinaryReader): SolveResponse;
}

export namespace SolveResponse {
    export type AsObject = {
        diff: string,
        error: string,
    }
}

export class ErrorResponse extends jspb.Message { 
    getError(): string;
    setError(value: string): ErrorResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ErrorResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ErrorResponse): ErrorResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ErrorResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ErrorResponse;
    static deserializeBinaryFromReader(message: ErrorResponse, reader: jspb.BinaryReader): ErrorResponse;
}

export namespace ErrorResponse {
    export type AsObject = {
        error: string,
    }
}

export class GetTokenRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): GetTokenRequest;
    getPassword(): string;
    setPassword(value: string): GetTokenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetTokenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetTokenRequest): GetTokenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetTokenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetTokenRequest;
    static deserializeBinaryFromReader(message: GetTokenRequest, reader: jspb.BinaryReader): GetTokenRequest;
}

export namespace GetTokenRequest {
    export type AsObject = {
        username: string,
        password: string,
    }
}

export class GetTokenResponse extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetTokenResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetTokenResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetTokenResponse): GetTokenResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetTokenResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetTokenResponse;
    static deserializeBinaryFromReader(message: GetTokenResponse, reader: jspb.BinaryReader): GetTokenResponse;
}

export namespace GetTokenResponse {
    export type AsObject = {
        token: string,
    }
}

export enum DiagramType {
    UNSPECIFIED = 0,
    CLASS = 1,
    SEQUENCE = 2,
}

export enum MemgerShip {
    BRONZE = 0,
    SILVER = 1,
    GOLD = 2,
}

export enum QuizStatus {
    DRAFT = 0,
    REVIEW = 1,
    PUBLIC = 2,
    REPORT = 3,
}
