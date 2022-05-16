// package: umlquiz
// file: umlquiz.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as umlquiz_pb from "./umlquiz_pb";
import * as google_type_datetime_pb from "./google/type/datetime_pb";

interface IUMLQuizLoginServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getToken: IUMLQuizLoginServiceService_IGetToken;
}

interface IUMLQuizLoginServiceService_IGetToken extends grpc.MethodDefinition<umlquiz_pb.GetTokenRequest, umlquiz_pb.GetTokenResponse> {
    path: "/umlquiz.UMLQuizLoginService/GetToken";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.GetTokenRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.GetTokenRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.GetTokenResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.GetTokenResponse>;
}

export const UMLQuizLoginServiceService: IUMLQuizLoginServiceService;

export interface IUMLQuizLoginServiceServer {
    getToken: grpc.handleUnaryCall<umlquiz_pb.GetTokenRequest, umlquiz_pb.GetTokenResponse>;
}

export interface IUMLQuizLoginServiceClient {
    getToken(request: umlquiz_pb.GetTokenRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.GetTokenResponse) => void): grpc.ClientUnaryCall;
    getToken(request: umlquiz_pb.GetTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.GetTokenResponse) => void): grpc.ClientUnaryCall;
    getToken(request: umlquiz_pb.GetTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.GetTokenResponse) => void): grpc.ClientUnaryCall;
}

export class UMLQuizLoginServiceClient extends grpc.Client implements IUMLQuizLoginServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getToken(request: umlquiz_pb.GetTokenRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.GetTokenResponse) => void): grpc.ClientUnaryCall;
    public getToken(request: umlquiz_pb.GetTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.GetTokenResponse) => void): grpc.ClientUnaryCall;
    public getToken(request: umlquiz_pb.GetTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.GetTokenResponse) => void): grpc.ClientUnaryCall;
}

interface IUMLQuizUserServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    addUser: IUMLQuizUserServiceService_IAddUser;
    updateUser: IUMLQuizUserServiceService_IUpdateUser;
    findUser: IUMLQuizUserServiceService_IFindUser;
    deleteUser: IUMLQuizUserServiceService_IDeleteUser;
}

interface IUMLQuizUserServiceService_IAddUser extends grpc.MethodDefinition<umlquiz_pb.AddUserRequest, umlquiz_pb.UserResponse> {
    path: "/umlquiz.UMLQuizUserService/AddUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.AddUserRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.AddUserRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.UserResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.UserResponse>;
}
interface IUMLQuizUserServiceService_IUpdateUser extends grpc.MethodDefinition<umlquiz_pb.UpdateUserRequest, umlquiz_pb.UserResponse> {
    path: "/umlquiz.UMLQuizUserService/UpdateUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.UpdateUserRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.UpdateUserRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.UserResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.UserResponse>;
}
interface IUMLQuizUserServiceService_IFindUser extends grpc.MethodDefinition<umlquiz_pb.UserRequest, umlquiz_pb.UserResponse> {
    path: "/umlquiz.UMLQuizUserService/FindUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.UserRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.UserRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.UserResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.UserResponse>;
}
interface IUMLQuizUserServiceService_IDeleteUser extends grpc.MethodDefinition<umlquiz_pb.UserRequest, umlquiz_pb.ErrorResponse> {
    path: "/umlquiz.UMLQuizUserService/DeleteUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.UserRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.UserRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.ErrorResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.ErrorResponse>;
}

export const UMLQuizUserServiceService: IUMLQuizUserServiceService;

export interface IUMLQuizUserServiceServer {
    addUser: grpc.handleUnaryCall<umlquiz_pb.AddUserRequest, umlquiz_pb.UserResponse>;
    updateUser: grpc.handleUnaryCall<umlquiz_pb.UpdateUserRequest, umlquiz_pb.UserResponse>;
    findUser: grpc.handleUnaryCall<umlquiz_pb.UserRequest, umlquiz_pb.UserResponse>;
    deleteUser: grpc.handleUnaryCall<umlquiz_pb.UserRequest, umlquiz_pb.ErrorResponse>;
}

export interface IUMLQuizUserServiceClient {
    addUser(request: umlquiz_pb.AddUserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    addUser(request: umlquiz_pb.AddUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    addUser(request: umlquiz_pb.AddUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    updateUser(request: umlquiz_pb.UpdateUserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    updateUser(request: umlquiz_pb.UpdateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    updateUser(request: umlquiz_pb.UpdateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    findUser(request: umlquiz_pb.UserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    findUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    findUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    deleteUser(request: umlquiz_pb.UserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    deleteUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    deleteUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
}

export class UMLQuizUserServiceClient extends grpc.Client implements IUMLQuizUserServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public addUser(request: umlquiz_pb.AddUserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public addUser(request: umlquiz_pb.AddUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public addUser(request: umlquiz_pb.AddUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public updateUser(request: umlquiz_pb.UpdateUserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public updateUser(request: umlquiz_pb.UpdateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public updateUser(request: umlquiz_pb.UpdateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public findUser(request: umlquiz_pb.UserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public findUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public findUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public deleteUser(request: umlquiz_pb.UserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public deleteUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public deleteUser(request: umlquiz_pb.UserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
}

interface IUMLQuizQuizServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    addQuiz: IUMLQuizQuizServiceService_IAddQuiz;
    findQuiz: IUMLQuizQuizServiceService_IFindQuiz;
    updateQuiz: IUMLQuizQuizServiceService_IUpdateQuiz;
    deleteQuiz: IUMLQuizQuizServiceService_IDeleteQuiz;
    listQuizzesAll: IUMLQuizQuizServiceService_IListQuizzesAll;
    listQuizzesByUser: IUMLQuizQuizServiceService_IListQuizzesByUser;
    solveQuiz: IUMLQuizQuizServiceService_ISolveQuiz;
    likeQuiz: IUMLQuizQuizServiceService_ILikeQuiz;
}

interface IUMLQuizQuizServiceService_IAddQuiz extends grpc.MethodDefinition<umlquiz_pb.AddQuizRequest, umlquiz_pb.QuizResponse> {
    path: "/umlquiz.UMLQuizQuizService/AddQuiz";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.AddQuizRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.AddQuizRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.QuizResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.QuizResponse>;
}
interface IUMLQuizQuizServiceService_IFindQuiz extends grpc.MethodDefinition<umlquiz_pb.FindQuizRequest, umlquiz_pb.QuizResponse> {
    path: "/umlquiz.UMLQuizQuizService/FindQuiz";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.FindQuizRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.FindQuizRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.QuizResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.QuizResponse>;
}
interface IUMLQuizQuizServiceService_IUpdateQuiz extends grpc.MethodDefinition<umlquiz_pb.UpdateQuizRequest, umlquiz_pb.QuizResponse> {
    path: "/umlquiz.UMLQuizQuizService/UpdateQuiz";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.UpdateQuizRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.UpdateQuizRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.QuizResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.QuizResponse>;
}
interface IUMLQuizQuizServiceService_IDeleteQuiz extends grpc.MethodDefinition<umlquiz_pb.DeleteQuizRequest, umlquiz_pb.ErrorResponse> {
    path: "/umlquiz.UMLQuizQuizService/DeleteQuiz";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.DeleteQuizRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.DeleteQuizRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.ErrorResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.ErrorResponse>;
}
interface IUMLQuizQuizServiceService_IListQuizzesAll extends grpc.MethodDefinition<umlquiz_pb.ListQuizzesAllRequest, umlquiz_pb.QuizzesResponse> {
    path: "/umlquiz.UMLQuizQuizService/ListQuizzesAll";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.ListQuizzesAllRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.ListQuizzesAllRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.QuizzesResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.QuizzesResponse>;
}
interface IUMLQuizQuizServiceService_IListQuizzesByUser extends grpc.MethodDefinition<umlquiz_pb.ListQuizzesByUserRequest, umlquiz_pb.QuizzesResponse> {
    path: "/umlquiz.UMLQuizQuizService/ListQuizzesByUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.ListQuizzesByUserRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.ListQuizzesByUserRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.QuizzesResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.QuizzesResponse>;
}
interface IUMLQuizQuizServiceService_ISolveQuiz extends grpc.MethodDefinition<umlquiz_pb.SolveQuizRequest, umlquiz_pb.SolveResponse> {
    path: "/umlquiz.UMLQuizQuizService/SolveQuiz";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.SolveQuizRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.SolveQuizRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.SolveResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.SolveResponse>;
}
interface IUMLQuizQuizServiceService_ILikeQuiz extends grpc.MethodDefinition<umlquiz_pb.LikeQuizRequest, umlquiz_pb.ErrorResponse> {
    path: "/umlquiz.UMLQuizQuizService/LikeQuiz";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.LikeQuizRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.LikeQuizRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.ErrorResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.ErrorResponse>;
}

export const UMLQuizQuizServiceService: IUMLQuizQuizServiceService;

export interface IUMLQuizQuizServiceServer {
    addQuiz: grpc.handleUnaryCall<umlquiz_pb.AddQuizRequest, umlquiz_pb.QuizResponse>;
    findQuiz: grpc.handleUnaryCall<umlquiz_pb.FindQuizRequest, umlquiz_pb.QuizResponse>;
    updateQuiz: grpc.handleUnaryCall<umlquiz_pb.UpdateQuizRequest, umlquiz_pb.QuizResponse>;
    deleteQuiz: grpc.handleUnaryCall<umlquiz_pb.DeleteQuizRequest, umlquiz_pb.ErrorResponse>;
    listQuizzesAll: grpc.handleUnaryCall<umlquiz_pb.ListQuizzesAllRequest, umlquiz_pb.QuizzesResponse>;
    listQuizzesByUser: grpc.handleUnaryCall<umlquiz_pb.ListQuizzesByUserRequest, umlquiz_pb.QuizzesResponse>;
    solveQuiz: grpc.handleUnaryCall<umlquiz_pb.SolveQuizRequest, umlquiz_pb.SolveResponse>;
    likeQuiz: grpc.handleUnaryCall<umlquiz_pb.LikeQuizRequest, umlquiz_pb.ErrorResponse>;
}

export interface IUMLQuizQuizServiceClient {
    addQuiz(request: umlquiz_pb.AddQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    addQuiz(request: umlquiz_pb.AddQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    addQuiz(request: umlquiz_pb.AddQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    findQuiz(request: umlquiz_pb.FindQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    findQuiz(request: umlquiz_pb.FindQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    findQuiz(request: umlquiz_pb.FindQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    updateQuiz(request: umlquiz_pb.UpdateQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    updateQuiz(request: umlquiz_pb.UpdateQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    updateQuiz(request: umlquiz_pb.UpdateQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    deleteQuiz(request: umlquiz_pb.DeleteQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    deleteQuiz(request: umlquiz_pb.DeleteQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    deleteQuiz(request: umlquiz_pb.DeleteQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    listQuizzesAll(request: umlquiz_pb.ListQuizzesAllRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    listQuizzesAll(request: umlquiz_pb.ListQuizzesAllRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    listQuizzesAll(request: umlquiz_pb.ListQuizzesAllRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    listQuizzesByUser(request: umlquiz_pb.ListQuizzesByUserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    listQuizzesByUser(request: umlquiz_pb.ListQuizzesByUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    listQuizzesByUser(request: umlquiz_pb.ListQuizzesByUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    solveQuiz(request: umlquiz_pb.SolveQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.SolveResponse) => void): grpc.ClientUnaryCall;
    solveQuiz(request: umlquiz_pb.SolveQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.SolveResponse) => void): grpc.ClientUnaryCall;
    solveQuiz(request: umlquiz_pb.SolveQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.SolveResponse) => void): grpc.ClientUnaryCall;
    likeQuiz(request: umlquiz_pb.LikeQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    likeQuiz(request: umlquiz_pb.LikeQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    likeQuiz(request: umlquiz_pb.LikeQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
}

export class UMLQuizQuizServiceClient extends grpc.Client implements IUMLQuizQuizServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public addQuiz(request: umlquiz_pb.AddQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public addQuiz(request: umlquiz_pb.AddQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public addQuiz(request: umlquiz_pb.AddQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public findQuiz(request: umlquiz_pb.FindQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public findQuiz(request: umlquiz_pb.FindQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public findQuiz(request: umlquiz_pb.FindQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public updateQuiz(request: umlquiz_pb.UpdateQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public updateQuiz(request: umlquiz_pb.UpdateQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public updateQuiz(request: umlquiz_pb.UpdateQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizResponse) => void): grpc.ClientUnaryCall;
    public deleteQuiz(request: umlquiz_pb.DeleteQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public deleteQuiz(request: umlquiz_pb.DeleteQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public deleteQuiz(request: umlquiz_pb.DeleteQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public listQuizzesAll(request: umlquiz_pb.ListQuizzesAllRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    public listQuizzesAll(request: umlquiz_pb.ListQuizzesAllRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    public listQuizzesAll(request: umlquiz_pb.ListQuizzesAllRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    public listQuizzesByUser(request: umlquiz_pb.ListQuizzesByUserRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    public listQuizzesByUser(request: umlquiz_pb.ListQuizzesByUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    public listQuizzesByUser(request: umlquiz_pb.ListQuizzesByUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.QuizzesResponse) => void): grpc.ClientUnaryCall;
    public solveQuiz(request: umlquiz_pb.SolveQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.SolveResponse) => void): grpc.ClientUnaryCall;
    public solveQuiz(request: umlquiz_pb.SolveQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.SolveResponse) => void): grpc.ClientUnaryCall;
    public solveQuiz(request: umlquiz_pb.SolveQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.SolveResponse) => void): grpc.ClientUnaryCall;
    public likeQuiz(request: umlquiz_pb.LikeQuizRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public likeQuiz(request: umlquiz_pb.LikeQuizRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public likeQuiz(request: umlquiz_pb.LikeQuizRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
}

interface IUMLQuizReportServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    addReport: IUMLQuizReportServiceService_IAddReport;
    findReports: IUMLQuizReportServiceService_IFindReports;
    updateReport: IUMLQuizReportServiceService_IUpdateReport;
    deleteReport: IUMLQuizReportServiceService_IDeleteReport;
}

interface IUMLQuizReportServiceService_IAddReport extends grpc.MethodDefinition<umlquiz_pb.AddReportRequest, umlquiz_pb.ReportResponse> {
    path: "/umlquiz.UMLQuizReportService/AddReport";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.AddReportRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.AddReportRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.ReportResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.ReportResponse>;
}
interface IUMLQuizReportServiceService_IFindReports extends grpc.MethodDefinition<umlquiz_pb.FindReportsRequest, umlquiz_pb.ReportsResponse> {
    path: "/umlquiz.UMLQuizReportService/FindReports";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.FindReportsRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.FindReportsRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.ReportsResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.ReportsResponse>;
}
interface IUMLQuizReportServiceService_IUpdateReport extends grpc.MethodDefinition<umlquiz_pb.UpdateReportRequest, umlquiz_pb.ReportResponse> {
    path: "/umlquiz.UMLQuizReportService/UpdateReport";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.UpdateReportRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.UpdateReportRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.ReportResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.ReportResponse>;
}
interface IUMLQuizReportServiceService_IDeleteReport extends grpc.MethodDefinition<umlquiz_pb.DeleteReportRequest, umlquiz_pb.ErrorResponse> {
    path: "/umlquiz.UMLQuizReportService/DeleteReport";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<umlquiz_pb.DeleteReportRequest>;
    requestDeserialize: grpc.deserialize<umlquiz_pb.DeleteReportRequest>;
    responseSerialize: grpc.serialize<umlquiz_pb.ErrorResponse>;
    responseDeserialize: grpc.deserialize<umlquiz_pb.ErrorResponse>;
}

export const UMLQuizReportServiceService: IUMLQuizReportServiceService;

export interface IUMLQuizReportServiceServer {
    addReport: grpc.handleUnaryCall<umlquiz_pb.AddReportRequest, umlquiz_pb.ReportResponse>;
    findReports: grpc.handleUnaryCall<umlquiz_pb.FindReportsRequest, umlquiz_pb.ReportsResponse>;
    updateReport: grpc.handleUnaryCall<umlquiz_pb.UpdateReportRequest, umlquiz_pb.ReportResponse>;
    deleteReport: grpc.handleUnaryCall<umlquiz_pb.DeleteReportRequest, umlquiz_pb.ErrorResponse>;
}

export interface IUMLQuizReportServiceClient {
    addReport(request: umlquiz_pb.AddReportRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    addReport(request: umlquiz_pb.AddReportRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    addReport(request: umlquiz_pb.AddReportRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    findReports(request: umlquiz_pb.FindReportsRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportsResponse) => void): grpc.ClientUnaryCall;
    findReports(request: umlquiz_pb.FindReportsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportsResponse) => void): grpc.ClientUnaryCall;
    findReports(request: umlquiz_pb.FindReportsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportsResponse) => void): grpc.ClientUnaryCall;
    updateReport(request: umlquiz_pb.UpdateReportRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    updateReport(request: umlquiz_pb.UpdateReportRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    updateReport(request: umlquiz_pb.UpdateReportRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    deleteReport(request: umlquiz_pb.DeleteReportRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    deleteReport(request: umlquiz_pb.DeleteReportRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    deleteReport(request: umlquiz_pb.DeleteReportRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
}

export class UMLQuizReportServiceClient extends grpc.Client implements IUMLQuizReportServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public addReport(request: umlquiz_pb.AddReportRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    public addReport(request: umlquiz_pb.AddReportRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    public addReport(request: umlquiz_pb.AddReportRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    public findReports(request: umlquiz_pb.FindReportsRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportsResponse) => void): grpc.ClientUnaryCall;
    public findReports(request: umlquiz_pb.FindReportsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportsResponse) => void): grpc.ClientUnaryCall;
    public findReports(request: umlquiz_pb.FindReportsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportsResponse) => void): grpc.ClientUnaryCall;
    public updateReport(request: umlquiz_pb.UpdateReportRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    public updateReport(request: umlquiz_pb.UpdateReportRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    public updateReport(request: umlquiz_pb.UpdateReportRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ReportResponse) => void): grpc.ClientUnaryCall;
    public deleteReport(request: umlquiz_pb.DeleteReportRequest, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public deleteReport(request: umlquiz_pb.DeleteReportRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
    public deleteReport(request: umlquiz_pb.DeleteReportRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: umlquiz_pb.ErrorResponse) => void): grpc.ClientUnaryCall;
}
