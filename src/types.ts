export type SuccessResult<Data> = {
  isSuccess: true;
  data: Data;
};

export type ErrorResult = {
  isSuccess: false;
  error: Error;
};

export type Result<Data> = SuccessResult<Data> | ErrorResult;

export type Weight = {
  weight: number;
  timestamp: string;
};
