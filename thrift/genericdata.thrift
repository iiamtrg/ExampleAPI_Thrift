namespace go myGeneric

enum TErrorCode {
    SUCCESS = 200,
    ITEM_NOT_EXISTED = 404,
    ITEM_ALREADY_EXISTED = 422,
    UNKNOWN_EXCEPTION = 500,
}

struct TDate{
    1: required i32 year,
    2: required i32 month,
    3: required i32 day,
}

struct TPerson {
    1: required string personId,
    2: optional string personName,
    3: optional TDate birthDate,
    4: optional string personAddress,
    5: optional TTeam team
}

struct TTeam {
    1: required string teamId,
    2: optional string teamName,
    3: optional string teamAddress,
    4: optional list<TPerson> persons
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

    void putItemPerson(1: string bsKey, 2: TPerson item),

    TTeamResult getItemTeam(1: string bsKey, 2: string rootID),

    TTeamSetResult getItemsTeam(1: string bsKey),

    void putItemTeam(1: string bsKey, 2: TTeam item),

    void removeItem(1: string bsKey, 2: string rootID),

    void removeAll(1: string bsKey)

}