def remove_exclamation_marks(s):
    s_l = list(s)
    s_r = ""
    for i in s_l:
        if i!="!":
            s_r+=i
    return s_r
        