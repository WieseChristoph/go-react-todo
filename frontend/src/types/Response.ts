type Response<T = undefined> = {
  success: boolean;
  message?: string;
  data: T;
};

export default Response;
