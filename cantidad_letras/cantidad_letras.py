import string


def open_file():
    fname = input('Type the filename: ')
    fhandler = open(fname, encoding='utf-8')
    return fhandler


def count_characters(filehandler, text):
    # print('------Printin filehandler: ', filehandler)

    char_count = dict()
    for line in filehandler:
        line = line.strip().lower()
        line = line.translate(line.maketrans('', '', string.punctuation))
        # print('--After doing strip each char, Character: ', line)
        text = text + line
    print('')
    print('____________________________________________')
    print('The text after concatenate lines: ', text)
    print('|')
    print('|___Type of the text variable: ', type(text))
    print('____________________________________________')
    # tratar de no hacer un for anidado aca, eso es lo que hay que mejorar de este codio

    for character in text:
        # print('')
        # print('Char in text:', character)
        if character.isalpha():
            if character in char_count:
                char_count[str(character)] += 1
            else:
                char_count[str(character)] = 1
            # char_count[character] = char_count.get(character)
        else:
            continue
    return char_count


def order_by_decreasing(counter):
    inverse_counter_lst = list()
    for element in counter:
        inverse_counter_lst.append((counter[element], element))

        inverse_counter_lst.sort(reverse=True)

        for number, element in inverse_counter_lst:
            print(f'{element} -> {number}')


first_text = ""
# order_by_decreasing(count_characters(open_file(), first_text))
print(count_characters(open_file(), first_text))
