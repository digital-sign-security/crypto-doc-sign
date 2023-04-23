CREATE TABLE public.user(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(128) UNIQUE NOT NULL,
    email VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(256) UNIQUE NOT NULL
);

CREATE TABLE public.jwttoken(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token VARCHAR(128) UNIQUE NOT NULL,
    is_alive BOOLEAN NOT NULL DEFAULT false,
    user_id UUID NOT NULL,
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.user(id)
);

CREATE TABLE public.keys(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    public_key TEXT NOT NULL,
    private_key TEXT NOT NULL,
    is_alive BOOLEAN NOT NULL DEFAULT false,
    user_id UUID NOT NULL,
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.user(id)
);

CREATE TABLE public.docs(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hash_ds TEXT NOT NULL,
    decrypted_text TEXT NOT NULL,
    sender_user_id UUID NOT NULL,
    recipient_user_id UUID NOT NULL,
    CONSTRAINT sender_user_fk FOREIGN KEY (sender_user_id) REFERENCES public.user(id),
    CONSTRAINT recipient_user_fk FOREIGN KEY (recipient_user_id) REFERENCES public.user(id)
);