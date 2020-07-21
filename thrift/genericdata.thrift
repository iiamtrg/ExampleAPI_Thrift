namespace go myGeneric

enum TErrorCode {
    SUCCESS = 200,
    ITEM_NOT_EXISTED = 404,
    ITEM_ALREADY_EXISTED = 422,
    UNKNOWN_EXCEPTION = 500,
}



struct TPerson {
    1: required string personId,
    2: optional string personName,
    3: optional string birthDate,
    4: optional string personAddress,
    5: optional string teamId
}

struct TTeam {
    1: required string teamId,
    2: optional string teamName,
    3: optional string teamAddress
}


struct TPersonResult{
    1: TErrorCode error,
    2: optional TPerson item
}

struct TPeronSetResult{
    1: TErrorCode error,
    2: optional list<TPerson> items
}

struct TTeamResult {
    1: TErrorCode error,
    2: optional TTeam item
}
struct TTeamSetResult{
    1: TErrorCode error,
    2: optional list<TTeam> items
}



service TGenericService {

    TPersonResult getItemPerson(1: string bsKey, 2: string rootID),

    TPeronSetResult getItemsPerson(1: string bsKey),

    TPeronSetResult getPersonsPagination(1: string bsKey, 2: i32 offset, 3: i32 limit)

    TPeronSetResult getPersonsOfTeam(1: string teamID),

    TPeronSetResult getPersonsOfTeamPagination(1: string teamID, 2: i32 offset, 3: i32 limit)

    TTeamResult getItemTeam(1: string bsKey, 2: string rootID),

    TTeamSetResult getItemsTeam(1: string bsKey),

    TTeamSetResult getTeamsPagination(1: string bsKey, 2: i32 offset, 3: i32 limit)

    TTeamResult getPersonIsTeam(1: string personId),

    void putItemPerson(1: string bsKey, 2: TPerson item),

    void putItemTeam(1: string bsKey, 2: TTeam item),

    void putPersonToTeam(1: string teamID, 2: string personId)

    bool itemIsExist(1: string bsKey, 2: string rootID),

    void removeItem(1: string bsKey, 2: string rootID),

}